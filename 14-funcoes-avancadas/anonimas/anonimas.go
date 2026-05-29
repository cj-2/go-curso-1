package main

import "fmt"

func main() {
	// Função anônima que se executa.
	func(texto string) {
		fmt.Println("Pa!", texto)
	}("O texto.") // Parâmetros.

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido => %s", texto)
	}("Jubileu")

	fmt.Println(retorno)
}
