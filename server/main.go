/*
gRPC Server
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sl "github.com/nleiva/lab-inventory/slack"
)

var (
	user     = getenv("USER")
	password = getenv("PASSWORD")
	address  = getenv("ADDRESS")
	database = getenv("DB")
	token    = getenv("SLACK_TOKEN")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Panicf("%s environment variable not set.", name)
	}
	return v
}

func main() {
	// Setup the DB
	target := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, address, database)
	db, err := sql.Open("mysql", target)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Validate connectivity to the DB.
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(fmt.Errorf("mysql: could not establish a good connection: %v", err))
	}

	// Initiate the Slack BOT
	ch := sl.Listen(db, token)
	log.Println("Slack BOT started")

	// Log the Slack actions (commands).
	for m := range ch {
		log.Printf("Executing: %v\n", m)
	}
}
