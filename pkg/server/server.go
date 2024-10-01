package server

import (
	"log"
	"net/http"
	"os"

	"github.com/tudemaha/toode/internal/shortener/controller"
)

func StartServer() {
	log.Println("INFO: StartServer: server is starting")

	http.Handle("/", controller.ShortenerSolverHandler())

	log.Println("INFO: StartServer: server started successfully")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
