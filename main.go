package main

import (
	"log"
	"net/http"

	"secure-login/router"
)

func main() {
	server := &http.Server{
		Addr:    ":8181",
		Handler: router.New(),
	}

	log.Println("server running on http://localhost:8181")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
