package s3

import (
	"log"
)

// Package variable
var auth *Auth

func init() {
	var err error
	auth, err = NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}
}
