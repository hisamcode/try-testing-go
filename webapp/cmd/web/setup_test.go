package main

import (
	"log"
	"os"
	"testing"

	"github.com/hisamcode/try-testing-go/webapp/pkg/db"
)

var app application

// will execute before actual test
func TestMain(m *testing.M) {
	pathToTemplates = "./../../../webapp/templates/"
	app.Session = getSession()
	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
