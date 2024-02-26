package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("sso/cmd/templates/index.html")
	tmpl.ExecuteTemplate(w, "index", nil)
}
