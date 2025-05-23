package serviceQuote

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (sq *serviceQuote) GetRandom(ctx context.Context) (*types.Quote, error) {
	return sq.repo.ServiceQuote.GetRandom(ctx)
}
