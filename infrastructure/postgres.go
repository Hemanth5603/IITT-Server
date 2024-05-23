package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var POSTGRES_DB *sql.DB
var POSTGRES_CONNECTION_STRING string

func init() {
	InitializePostgresSQL()
}

func InitializePostgresSQL() {
	var err error
	USER := "postgres"
	PASS := "admin"
	HOST := "localhost"
	DBNAME := "IITT"
	PORT := "5433"

	POSTGRES_CONNECTION_STRING = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DBNAME, PORT)

	POSTGRES_DB, err = sql.Open("pgx", POSTGRES_CONNECTION_STRING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to POSTGRES database: %v\n", err)
		os.Exit(1)
	}
	POSTGRES_DB.SetMaxIdleConns(10)

}
