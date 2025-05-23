package serviceQuote

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (sq *serviceQuote) GetAll(ctx context.Context, author string) ([]types.Quote, error) {
	return sq.repo.ServiceQuote.GetAll(ctx, author)
}
