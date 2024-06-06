package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func generateHash(word string) string {
	h := sha256.New()
	h.Write([]byte(os.Getenv("SECRET") + word))
	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString
}
