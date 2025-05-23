package psql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/molodoymaxim/service-quotes/internal/types"
)

func (t *psql) GetRandom(ctx context.Context) (*types.Quote, error) {
	var quote types.Quote
	query := `SELECT id, author, quote FROM quotes ORDER BY RANDOM() LIMIT 1`

	row := t.db.QueryRow(ctx, query)

	err := row.Scan(&quote.ID, &quote.Author, &quote.Quote)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no random quote found")
		}
		return nil, err
	}
	return &quote, nil
}
