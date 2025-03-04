package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var connection *pgx.Conn

func ConnectDB() {
	var err error
	connection, err = pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
}

func CloseDB() {
	if connection != nil {
		connection.Close(context.Background())
	}
}
