package database

import (
	"context"
	"fmt"
	"github.com/molodoymaxim/service-quotes/internal/system/database/psql"
	"github.com/molodoymaxim/service-quotes/internal/types"
	"time"
)

type DataBase struct {
	PSQL psql.Postgres
}

func New(cfgConnPostgres *types.ConfigConnPostgres, cfgPostgres *types.ConfigPostgres) (*DataBase, error) {
	сtxConn, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Подключение к Postgres
	psql := psql.New(
		cfgPostgres.PostgresUser,
		cfgPostgres.PostgresPassword,
		cfgPostgres.PostgresHost,
		cfgPostgres.PostgresDBName,
		cfgPostgres.PostgresPort,
		cfgPostgres.CfgDBTimeout,
		cfgPostgres.CfgDBQueryTimeout,
	)

	// Настройка пула
	err := psql.NewPoolConfig(
		cfgPostgres.CfgDBMaxConn,
		time.Duration(cfgPostgres.CfgDBConnIdleTime)*time.Second,
		time.Duration(cfgPostgres.CfgDBConnLifeTime)*time.Second,
	)
	if err != nil {
		return nil, fmt.Errorf("postgres new pool config: %v", err)
	}

	// Подключаем пулл подключений Postgres
	if err = psql.ConnectionPool(сtxConn); err != nil {
		return nil, fmt.Errorf("postgres connection pool: %v", err)
	}

	return &DataBase{
		PSQL: psql,
	}, nil
}
