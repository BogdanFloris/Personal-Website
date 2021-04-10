package accessor

import (
	"bogdanfloris-com/pkg/logging"
	"context"
	"github.com/jackc/pgx/v4"
	"os"
)

type BfdbAccessor struct {
	conn *pgx.Conn
}

func GetAccessor() *BfdbAccessor {
	bfdbConn, err := pgx.Connect(context.Background(), os.Getenv("BF_DATABASE_URL"))
	if err != nil {
		logging.ErrorLogger.Println("Cannot connect to postgres database bfdb.")
	}
	logging.InfoLogger.Println("Connection to bfdb database acquired.")

	accessor := &BfdbAccessor{bfdbConn}
	return accessor
}

func (accessor *BfdbAccessor) Close() {
	err := accessor.conn.Close(context.Background())
	if err != nil {
		logging.ErrorLogger.Println("Failed to close connection to bfdb database.")
	}
	logging.InfoLogger.Println("Connection to bfdb database closed.")
}

func (accessor *BfdbAccessor) Test() {
	var tables string
	err := accessor.conn.QueryRow(context.Background(), "SELECT current_database()").Scan(&tables)
	if err != nil {
		logging.ErrorLogger.Println(err)
	}
	logging.InfoLogger.Println(tables)
}
