package main

import (
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	// request é um ponteiro
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Carlos José", "carlos.jose@gmail.com"}
		templates.ExecuteTemplate(w, "index.html", u)
	})

	// Iniciando um servidor http com um pacote nativo:
	log.Fatal(http.ListenAndServe(":5000", nil))
}
