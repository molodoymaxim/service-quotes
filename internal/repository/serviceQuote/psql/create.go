package psql

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (t *psql) Create(ctx context.Context, q *types.Quote) error {
	query := `INSERT INTO quotes (author, quote) VALUES ($1, $2) RETURNING id`
	return t.db.QueryRow(ctx, query, q.Author, q.Text).Scan(&q.ID)
}
