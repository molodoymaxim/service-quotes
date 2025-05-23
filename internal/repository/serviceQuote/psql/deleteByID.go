package psql

import "context"

func (t *psql) DeleteByID(ctx context.Context, id int64) error {
	query := "DELETE FROM quotes WHERE id = $1"
	_, err := t.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
