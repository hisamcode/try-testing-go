package main

import (
	"os"
	"testing"

	"github.com/hisamcode/try-testing-go/webapp/pkg/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2OTUwMDUyMDYsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.2CfhfZ0D2ZGI9dCQDTP2YAtRtZ34KUpsP6GD-lAz7Kw"
var hs384Token = "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2OTU2MjgwMTUsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.Tcuc4c6rQbReaUfuCn-LuX5VdPR5OpYXVOlJZgdQIvjLhi4CwpazpLFOnCkMMdJf"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160"

	code := m.Run()

	os.Exit(code)
}
