package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Connection *pgx.Conn // Fixed: Exported variable so other packages can access it

func ConnectDB() {
	var err error
	Connection, err = pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1) // Ensures the program exits if the DB connection fails
	}
}

func CloseDB() {
	if Connection != nil {
		Connection.Close(context.Background())
	}
}
