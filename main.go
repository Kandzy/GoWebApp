package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title       string
	Message     string
	Description string
}

func Home(writer http.ResponseWriter, reader *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	data := PageData{
		Title:       "Home Page",
		Message:     "First message",
		Description: "Easy description",
	}
	tpl.Execute(writer, data)
}

func main() {
	log.Println("Starting Server ...")

	http.HandleFunc("/", Home)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
