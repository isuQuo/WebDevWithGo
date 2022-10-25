package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[3], os.Args[2])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed password: %s\n", string(hashedPassword))
}

func compare(hashedPassword string, plaintext string) {
	if ok := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintext)); ok != nil {
		fmt.Println("Invalid match")
		return
	}

	fmt.Println("Valid match")
}
