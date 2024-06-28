package handlers

import (
	"GoWebServer/src/pkg/handlers/homepage"
	"net/http"
)

func InitiateHandlers() {
	http.HandleFunc("/", homepage.Home)
}
