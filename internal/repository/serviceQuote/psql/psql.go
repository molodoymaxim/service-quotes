package psql

import (
	"context"
	p "github.com/molodoymaxim/service-quotes/internal/system/database/psql"

	"github.com/molodoymaxim/service-quotes/internal/types"
)

type PSQL interface {
	Create(ctx context.Context, q *types.Quote) error
	GetAll(ctx context.Context, author string) ([]types.Quote, error)
	GetRandom(ctx context.Context) (*types.Quote, error)
	DeleteByID(ctx context.Context, id int64) error
}

type psql struct {
	db p.Postgres
}

func New(db p.Postgres) PSQL {
	return &psql{
		db: db,
	}
}
