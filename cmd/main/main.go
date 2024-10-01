package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tudemaha/toode/pkg/mongodb"
	"github.com/tudemaha/toode/pkg/server"
)

func main() {
	mongodb.StartConnection()
	server.StartServer()
}
