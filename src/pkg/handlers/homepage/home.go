package homepage

import (
	"html/template"
	"net/http"
	"path/filepath"
	"GoWebServer/src/pkg/helpers/globalpath"
)

type PageData struct {
	Title       string
	Message     string
	Description string
}

func Home(writer http.ResponseWriter, reader *http.Request) {
	tplPath := "assets/templates/home.html"
	fullPath := filepath.Join(globalpath.GetExecutableLocation(), tplPath)

	tpl := template.Must(template.ParseFiles(fullPath))
	data := PageData{
		Title:       "Home Page",
		Message:     "First message",
		Description: "Easy description",
	}
	tpl.Execute(writer, data)
}
