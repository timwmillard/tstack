package main

import (
	"app/auth"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: hash <password>\n")
		os.Exit(1)
	}

	password := os.Args[1]
	hash, err := auth.HashPassword(password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] hashing password: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(hash)
}
