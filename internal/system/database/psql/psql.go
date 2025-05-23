package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/url"
	"runtime"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Postgres interface {
	NewPoolConfig(maxConn int, connIdleTime, connLifeTime time.Duration) error    // Создание конфигурации пула
	ConnectionPool(ctx context.Context) error                                     // Подключаемся с помощью пула к Postgres
	GetSQL(sqlFunc func(db *sql.DB) error) error                                  // Выполнение функции от имени драйвера sql.DB
	Ping(ctx context.Context) error                                               // Проверяем соединение
	Close()                                                                       // Закрытие соединения
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)         // Query запрос к БД
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row                // QueryRow запрос к БД
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) // Exec запрос к БД
}

type postgres struct {
	conn         *pgxpool.Pool
	connStr      string
	poolConfig   *pgxpool.Config ``
	queryTimeout time.Duration
}

func New(user, pass, host, dbName string, port, timeout, queryTimeout int) Postgres {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(user),
		url.QueryEscape(pass),
		host,
		port,
		dbName,
		timeout)

	return &postgres{
		connStr:      connStr,
		queryTimeout: time.Duration(queryTimeout) * time.Second,
	}
}

// Создание конфигурации пула
func (d *postgres) NewPoolConfig(maxConn int, connIdleTime, connLifeTime time.Duration) error {
	// Создание конфигурации пула
	poolConfig, err := pgxpool.ParseConfig(d.connStr)
	if err != nil {
		return err
	}

	// Проверка
	cpu := runtime.NumCPU()
	if maxConn > cpu {
		maxConn = cpu
	}

	poolConfig.MaxConns = int32(maxConn)
	poolConfig.MaxConnIdleTime = connIdleTime
	poolConfig.MaxConnLifetime = connLifeTime
	d.poolConfig = poolConfig
	return nil
}

// Подключаемся с помощью пула к Postgres
func (d *postgres) ConnectionPool(ctx context.Context) error {
	conn, err := pgxpool.NewWithConfig(ctx, d.poolConfig)
	if err != nil {
		return err
	}
	d.conn = conn
	err = d.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Проверяем соединение
func (d *postgres) Ping(ctx context.Context) error {
	return d.conn.Ping(ctx)
}

// Закрытие соединения
func (d *postgres) Close() {
	d.conn.Close()
}

// Обработчик транзакций
func (d *postgres) Transact(ctxParent context.Context, txFunc func(ctx context.Context, tx pgx.Tx) error) (err error) {
	ctx, cancel := context.WithTimeout(ctxParent, d.queryTimeout)
	defer cancel()

	tx, err := d.conn.Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if errRoll := tx.Rollback(ctx); errRoll != nil {
				log.Printf("Failed rollback database TX: %v\n", errRoll)
			}
			panic(p)
		} else if err != nil {
			if errRoll := tx.Rollback(ctx); errRoll != nil {
				log.Printf("Failed rollback database TX: %v\n", errRoll)
			}
		} else {
			err = tx.Commit(ctx)
		}
	}()
	err = txFunc(ctx, tx)
	return err
}

// Выполнение функции от имени драйвера sql.DB
func (d *postgres) GetSQL(sqlFunc func(db *sql.DB) error) error {
	return sqlFunc(stdlib.OpenDBFromPool(d.conn))
}

func (d *postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return d.conn.Query(ctx, sql, args...)
}

func (d *postgres) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return d.conn.QueryRow(ctx, sql, args...)
}

func (d *postgres) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return d.conn.Exec(ctx, sql, args...)
}
