package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	router := NewRouter()

	log.Info("Starting server on port 8000")
	http.ListenAndServe(":8000", router)
}
