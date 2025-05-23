package psql

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (t *psql) GetAll(ctx context.Context, author string) ([]types.Quote, error) {
	var quotes []types.Quote
	query := "SELECT id, author, quote FROM quotes"
	args := []interface{}{}

	if author != "" {
		query += " WHERE author = $1"
		args = append(args, author)
	}

	rows, err := t.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var q types.Quote
		err := rows.Scan(&q.ID, &q.Author, &q.Text)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quotes, err
}
