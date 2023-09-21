package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/hisamcode/try-testing-go/webapp/pkg/data"
	"github.com/hisamcode/try-testing-go/webapp/pkg/repository"
	"github.com/hisamcode/try-testing-go/webapp/pkg/repository/dbrepo"
)

type application struct {
	Session *scs.SessionManager
	DB      repository.DatabaseRepo
	DSN     string
}

func main() {
	// for the session, for put actual type in the session
	gob.Register(data.User{})

	// set up an app config
	app := application{}
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection data source name")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// get a session manager
	app.Session = getSession()

	// get application route
	mux := app.routes()

	// print out a message
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
