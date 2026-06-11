package main

import (
	"crud-basico/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))
}
