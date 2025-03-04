package users

import (
	"context"
	"fmt"
)

func GetUserById(id string) (string, string, error) {
	var username, email string
	err := db.connection.QueryRow(context.Background(), "select username, email from users where id=$1", id).
		Scan(&username, &email)
	if err != nil {
		return "", "", fmt.Errorf("Query failed: %v", err)
	}
	return username, email, nil
}

// func GetUsers() ([]string, error) {
//     var response []string
//     err := db.connection.Query(context.Background(), "select * from users")
// 	if err != nil {
// 		return nil, fmt.Errorf("Query failed: %v", err)
// 	}
//     return
// }
