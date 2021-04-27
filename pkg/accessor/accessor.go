package accessor

import (
	"bogdanfloris-com/pkg/logging"
	"bogdanfloris-com/pkg/models"
	"bogdanfloris-com/pkg/utils"
	"context"
	uuid2 "github.com/google/uuid"
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

func (accessor *PgAccessor) AddUser(username string, password string) (string, error) {
	// Hash Password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}
	// Generate uuid
	uuid := uuid2.NewString()
	// Sql statement
	sqlStatement := `
INSERT INTO "user" (user_id, username, password_hash, created_on)
VALUES ($1, $2, $3, NOW())`
	_, err = accessor.conn.Exec(context.Background(), sqlStatement, uuid, username, hash)
	return uuid, err
}

func (accessor *PgAccessor) RemoveUser(username string) error {
	_, err := accessor.conn.Exec(context.Background(), "DELETE FROM \"user\" WHERE username=$1", username)
	return err
}

func (accessor *PgAccessor) User(username string) (models.User, error) {
	var user models.User
	err := accessor.conn.QueryRow(context.Background(), "SELECT * FROM \"user\" WHERE username=$1", username).Scan(
		&user.Uuid, &user.Username, &user.PasswordHash, &user.CreatedOn, &user.LastLogin)
	return user, err
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
