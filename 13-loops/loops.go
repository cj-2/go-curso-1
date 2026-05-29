package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// Parecido com o while
	for i < 3 {
		// time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	// "+= 2" somaria de 2 em 2
	for e := 0; e < 10; e++ {
		fmt.Println(e)
	}

	nomes := [3]string{"Carlos", "Carlinhos", "Carlão"}

	// Iteração de uma coleção
	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for i, letra := range "Teste" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":  "Carlos",
		"idade": "28",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	for {
		time.Sleep(time.Second)
		fmt.Println("Fim do loop infinito!")
		break
	}
}
