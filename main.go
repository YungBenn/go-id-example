package main

import (
	"crypto/rand"
	"fmt"

	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/oklog/ulid/v2"
)

func main() {
	// Nano ID
	nanoID, err := generateNanoID()
	if err != nil {
		panic(err)
	}

	// CUID
	cuid, err := generateCUID()
	if err != nil {
		panic(err)
	}

	// ULID
	ulid := generateULID()

	// UUID v4
	uuidv4 := generateUUIDv4()

	// UUID v7
	uuidv7, err := generateUUIDv7()
	if err != nil {
		panic(err)
	}

	fmt.Println("NanoID:", nanoID)
	fmt.Println("CUID:", cuid)
	fmt.Println("ULID:", ulid)
	fmt.Println("UUIDv4:", uuidv4)
	fmt.Println("UUIDv7:", uuidv7)
}

func generateUUIDv4() string {
	uuid := uuid.New()
	return uuid.String()
}

func generateUUIDv7() (string, error) {
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return uuidv7.String(), nil
}

func generateNanoID() (string, error) {
	nanoID, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz1234567890", 23)
	if err != nil {
		return "", nil
	}

	return nanoID, nil
}

func generateCUID() (string, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return "", err
	}

	return cuid, nil
}

func generateULID() string {
	ulid := ulid.Make()
	return ulid.String()
}
