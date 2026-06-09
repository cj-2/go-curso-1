package main

import (
	"log"
	"net/http"
)

func main() {
	// request é um ponteiro
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo."))
	})

	// Iniciando um servidor http com um pacote nativo:
	log.Fatal(http.ListenAndServe(":5000", nil))
}
