package main

import (
	"GoWebServer/src/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server ...")

	handlers.InitiateHandlers()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
