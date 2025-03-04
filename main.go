package main

import (
	"fmt"
	"log"

	"github.com/egramsdoescode/go-backend-test/config"
	"github.com/egramsdoescode/go-backend-test/db"
	"github.com/egramsdoescode/go-backend-test/users"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()
	defer db.CloseDB()

	username, email, err := users.GetUserById(3)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}

	fmt.Printf("User: %s, Email: %s\n", username, email)
}
