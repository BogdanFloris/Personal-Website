package accessor

import (
	"bogdanfloris-com/pkg/logging"
	"context"
	"github.com/jackc/pgx/v4"
	"os"
)

type DbAccessor interface {
	Close()
	DatabaseName() string
}

type PgAccessor struct {
	conn *pgx.Conn
}

func NewPgAccessor() *PgAccessor {
	databaseUrl := os.Getenv("BF_DATABASE_URL")
	if len(databaseUrl) == 0 {
		logging.ErrorLogger.Println("BF_DATABASE_URL environment variable not set")
		os.Exit(1)
	}
	pgConn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		logging.ErrorLogger.Println("Cannot connect to postgres database bfdb.")
	}
	logging.InfoLogger.Println("Connection to bfdb database acquired.")

	accessor := &PgAccessor{pgConn}
	return accessor
}

func (accessor *PgAccessor) Close() {
	err := accessor.conn.Close(context.Background())
	if err != nil {
		logging.ErrorLogger.Println("Failed to close connection to bfdb database.")
	}
	logging.InfoLogger.Println("Connection to bfdb database closed.")
}

func (accessor *PgAccessor) DatabaseName() string {
	var name string
	err := accessor.conn.QueryRow(context.Background(), "SELECT current_database()").Scan(&name)
	if err != nil {
		logging.ErrorLogger.Println(err)
	}
	return name
}
