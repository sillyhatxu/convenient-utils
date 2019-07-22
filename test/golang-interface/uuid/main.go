package main

import (
	"fmt"
	"github.com/oklog/ulid"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

func main() {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u3, err := uuid.FromString("be182b4f-8f28-4dbb-9dde-3dd5c1fc36bd")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u3)
	fmt.Println()
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3

	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3

	//JAVA:    SKU 01D53Q9HD6MWZS9C3RYSF8GEJ9
	//go.uuid: SKU F46AD3A9878C4CC594B7E168B5832B7A
	//ulid:    SKU 3560114515464375470
}
