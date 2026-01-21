package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("conn_string")
	return pgx.Connect(ctx, connString)

}

//"postgres://postgres:admin@localhost:5432/postgres"
