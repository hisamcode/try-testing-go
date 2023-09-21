package main

import (
	"os"
	"testing"

	"github.com/hisamcode/try-testing-go/webapp/pkg/repository/dbrepo"
)

var app application

// will execute before actual test
func TestMain(m *testing.M) {
	pathToTemplates = "./../../../webapp/templates/"

	app.Session = getSession()
	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
