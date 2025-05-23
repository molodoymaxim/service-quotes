package serviceQuote

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (sq *serviceQuote) Create(ctx context.Context, q *types.Quote) error {
	//if q.Author == "" || q.Text == "" {
	//	return errors.New("Author and Text cannot be empty")
	//}
	return sq.repo.ServiceQuote.Create(ctx, q)
}
