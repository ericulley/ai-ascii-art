/*
Copyright © 2024 Eric Culley <ericulley>
*/
package main

import (
	"log"

	"github.com/ericulley/ascii/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
