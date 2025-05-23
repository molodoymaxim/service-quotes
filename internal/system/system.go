package system

import (
	"fmt"
	"github.com/molodoymaxim/service-quotes/internal/system/database"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

type Systems struct {
	DB *database.DataBase
}

func New(cfgConnPostgres *types.ConfigConnPostgres, cfgPostgres *types.ConfigPostgres) (*Systems, error) {
	db, err := database.New(
		cfgConnPostgres,
		cfgPostgres,
	)
	if err != nil {
		return nil, fmt.Errorf("database: %v", err)
	}
	return &Systems{
		DB: db,
	}, nil
}
