package serviceQuote

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/repository"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

type QuoteService interface {
	Create(ctx context.Context, q *types.Quote) error
	GetAll(ctx context.Context, author string) ([]types.Quote, error)
	GetRandom(ctx context.Context) (*types.Quote, error)
	DeleteByID(ctx context.Context, id int64) error
}

type serviceQuote struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) QuoteService {
	return &serviceQuote{
		repo: repo,
	}
}
