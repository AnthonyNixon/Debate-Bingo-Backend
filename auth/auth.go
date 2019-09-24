package auth

import (
	"log"
	"os"
	"time"
)

var JWT_SIGNING_KEY []byte

const TOKEN_VALID_TIME = 24 * time.Hour

func Initialize() {
	log.Print("Initializing Authentication")
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("No Signing Key Present.")
	}

	JWT_SIGNING_KEY = []byte(signingKey)
	log.Print("done")
}