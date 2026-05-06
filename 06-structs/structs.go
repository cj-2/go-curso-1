package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logadouro string
	numero    uint16
}

func main() {
	// declaração struct
	var usuario1 usuario
	usuario1.nome = "Carlos"
	usuario1.idade = 28

	// inferença de tipo, ordem importa
	usuario2 := usuario{"Marcos", 12, endereco{"Rua Y", 10}}

	// inferença de tipo com props name
	usuario3 := usuario{idade: 12}

	fmt.Println(usuario1)
	fmt.Println(usuario2)
	fmt.Println(usuario3)
}
