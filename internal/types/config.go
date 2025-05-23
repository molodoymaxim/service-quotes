package types

type ConfigHTTP struct {
	Port        int `mapstructure:"HTTP_PORT"`
	TimeLifeCtx int `mapstructure:"HTTP_TIME_LIFE_CTX"`
}

// Конфигурация подключения к БД Postgres
type ConfigPostgres struct {
	PostgresHost      string `mapstructure:"POSTGRES_HOST"`
	PostgresPort      int    `mapstructure:"POSTGRES_PORT"`
	PostgresUser      string `mapstructure:"POSTGRES_USER"`
	PostgresPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDBName    string `mapstructure:"POSTGRES_DB"`
	CfgDBTimeout      int    `mapstructure:"DB_TIMEOUT_PS"`
	CfgDBMaxConn      int    `mapstructure:"DB_MAX_CONN_PS"`
	CfgDBConnIdleTime int    `mapstructure:"DB_CONN_IDLE_TIME_PS"`
	CfgDBConnLifeTime int    `mapstructure:"DB_CONN_LIFE_TIME_PS"`
	CfgDBQueryTimeout int    `mapstructure:"DB_QUERY_TIMEOUT_PS"`
}

// Конфигурация подключения к БД Postgres
type ConfigConnPostgres struct {
	CfgDBMaxConn      int `mapstructure:"DB_MAX_CONN"`
	CfgDBMaxConnIdle  int `mapstructure:"DB_MAX_CONN_IDLE"`
	CfgDBConnIdleTime int `mapstructure:"DB_CONN_IDLE_TIME"`
	CfgDBConnLifeTime int `mapstructure:"DB_CONN_LIFE_TIME"`
}
