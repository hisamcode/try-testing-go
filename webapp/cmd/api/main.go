package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hisamcode/try-testing-go/webapp/pkg/repository"
	"github.com/hisamcode/try-testing-go/webapp/pkg/repository/dbrepo"
)

const port = 8090

type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {
	app := application{}

	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for application, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection data source name")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "signing secret")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Printf("starting api on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
