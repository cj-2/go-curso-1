package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorro1 := cachorro{"Bob", "Pintcher", 1}

	// struct normal
	fmt.Println(cachorro1)

	cachorroJson, erro1 := json.Marshal(cachorro1)

	if erro1 != nil {
		log.Fatal("Erro ao converter JSON")
	}

	// Convertido em JSON (slice de bytes)
	fmt.Println(cachorroJson)

	// "formatado"
	fmt.Println(bytes.NewBuffer(cachorroJson))

	// utilizando um map

	cachorro2 := map[string]string{
		"nome": "Juca",
		"raca": "Pintibu",
	}

	c2Json, erro2 := json.Marshal(cachorro2)

	if erro2 != nil {
		log.Fatal("Erro ao converter JSON")
	}

	fmt.Println(bytes.NewBuffer(c2Json))
}
