package homecontroller

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		"views/templates/mytemlate.html",
		"views/home/index.html",
	))
}

func Index(reponse http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(reponse, "layout", nil)
}
