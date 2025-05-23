package serviceQuote

import "context"

func (sq *serviceQuote) DeleteByID(ctx context.Context, id int64) error {
	return sq.repo.ServiceQuote.DeleteByID(ctx, id)
}
