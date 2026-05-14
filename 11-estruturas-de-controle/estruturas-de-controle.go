package main

import "fmt"

func main() {
	numero := 10

	// parênteses são opcionais
	if numero > 15 {
		fmt.Println("Maior que 15!")
	} else if numero == 15 {
		fmt.Println("Igual que 15!")
	} else {
		fmt.Println("Menor que 15!")
	}

	/*
		if init, declaramos uma variável "interna" ao
		escopo do if e já "validamos" ela dentro do if.

		Ela não existe fora da estrutura de condição
	*/
	if outroNumero := numero + 1; outroNumero > 0 {
		fmt.Println("Outro número é maior que zero...")
	} else {
		fmt.Println(outroNumero)
	}
}
