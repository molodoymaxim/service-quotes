package repository

import (
	"context"
	"fmt"
	"github.com/molodoymaxim/service-quotes/internal/repository/serviceQuote"
	"github.com/molodoymaxim/service-quotes/internal/system"
)

type Repository struct {
	ServiceQuote serviceQuote.ServiceQuote
	syst         *system.Systems
}

func New(syst *system.Systems) *Repository {
	return &Repository{
		syst:         syst,
		ServiceQuote: serviceQuote.New(syst),
	}
}

func (r *Repository) Ping(ctx context.Context) error {
	if err := r.syst.DB.PSQL.Ping(ctx); err != nil {
		return fmt.Errorf("postgres: %v", err)
	}
	return nil
}
