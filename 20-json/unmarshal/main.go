package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Bob","raca":"Pintcher","idade":1}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal("Erro ao converter JSON")
	}

	fmt.Println(c)

	cachorroEmJson2 := `{"nome":"Maicon","raca":"Bull","idade":"1"}`
	var c2 map[string]string

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal("Erro ao converter JSON")
	}

	fmt.Println(c2)
}
