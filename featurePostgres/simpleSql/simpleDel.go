package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, tasksIds []int) error {
	sqlQuery := `
		DELETE FROM tasks
		WHERE id = ANY($1);
	`
	_, err := conn.Exec(ctx, sqlQuery, tasksIds)
	return err
}
