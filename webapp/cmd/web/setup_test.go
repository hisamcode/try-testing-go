package main

import (
	"os"
	"testing"
)

var app application

// will execute before actual test
func TestMain(m *testing.M) {
	pathToTemplates = "./../../../webapp/templates/"
	app.Session = getSession()

	os.Exit(m.Run())
}
