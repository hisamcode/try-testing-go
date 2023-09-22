package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type application struct {
	JWTSecret string
	Action    string
}

// This is used to generate a token, so that we can test our api. Run this with go run ./cmd/cli and copy
// the token that is printed out.
// go run ./cmd/cli -action=valid     // will produce a valid token
// go run ./cmd/cli -action=expired   // will produce an expired token

func main() {
	var app application
	flag.StringVar(&app.JWTSecret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
	flag.StringVar(&app.Action, "action", "valid", "action: valid|expired|HS384")
	flag.Parse()

	// generate a token
	token := jwt.New(jwt.SigningMethodHS256)
	if app.Action == "HS384" {
		token = jwt.New(jwt.SigningMethodHS384)

	}

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["sub"] = "1"
	claims["admin"] = true
	claims["aud"] = "example.com"
	claims["iss"] = "example.com"
	// leave this to 3 days, for easy manual testing
	if app.Action == "valid" || app.Action == "HS384" {
		expires := time.Now().UTC().Add(time.Hour * 72)
		claims["exp"] = expires.Unix()
	} else {
		expires := time.Now().UTC().Add(time.Hour * 100 * -1)
		claims["exp"] = expires.Unix()
	}

	// create the token as a slice of bytes
	if app.Action == "valid" || app.Action == "HS384" {
		fmt.Println("VALID Token:")
	} else {
		fmt.Println("EXPIRED Token:")
	}
	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	// print to console
	fmt.Println(string(signedAccessToken))
}
