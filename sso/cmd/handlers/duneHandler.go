package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func DuneHandler(w http.ResponseWriter, r *http.Request, tokenJWT string, tokenExpiration time.Time, imageURL string) {
	switch {
	case tokenJWT == "":
		http.Error(w, "Unautorized", http.StatusUnauthorized)
	case time.Now().After(tokenExpiration):
		http.Error(w, "Unautorized", http.StatusUnauthorized)
	default:
		image := imageURL
		tmpl, err := template.ParseFiles("sso/cmd/templates/header.html", "sso/cmd/templates/dune.html", "sso/cmd/templates/player.html")
		if err != nil {
			fmt.Println("Failed to connect to DuneHandler")
		}
		tmpl.ExecuteTemplate(w, "dune", image)
	}
}
