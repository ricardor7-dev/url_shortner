package server

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

const URLLength = 7

func generateShortURL() string {
    b := make([]byte, URLLength)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatalf("Unable to generate short URL: %v", err)
    }
    return base64.RawURLEncoding.EncodeToString(b)
}