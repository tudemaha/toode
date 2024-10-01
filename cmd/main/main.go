package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tudemaha/toode/pkg/mongodb"
	"github.com/tudemaha/toode/pkg/server"
)

func main() {
	if os.Getenv("LOAD_DOTENV") == "yes" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("ERROR: Load loading .env file")
		}
	}

	mongodb.StartConnection()
	server.StartServer()
}
