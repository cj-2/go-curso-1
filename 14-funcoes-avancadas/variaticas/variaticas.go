package main

import "fmt"

// numéros é um slice
func soma(numeros ...int) (total int) {
	total = 0

	for _, valor := range numeros {
		total += valor
	}

	return
}

// Mesclando parâmetros comuns e variáticos
func exemploBobo(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	fmt.Println(soma(10, 20, 30, 40))
	exemploBobo("Abacate", 1, 2, 3)
}
