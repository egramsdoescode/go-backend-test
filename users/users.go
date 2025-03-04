package users

import (
	"context"
	"fmt"

	"github.com/egramsdoescode/go-backend-test/db"
)

func GetUserById(id int) (string, string, error) {
	var username, email string
	err := db.Connection.QueryRow(context.Background(), "SELECT username, email FROM users WHERE id=$1", id).
		Scan(&username, &email)
	if err != nil {
		return "", "", fmt.Errorf("Query failed: %v", err)
	}
	return username, email, nil
}
