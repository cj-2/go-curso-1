package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	matricula uint16
	curso     string
	polo      string
}

func main() {
	p1 := pessoa{"Carlos", "Roberto", 28, 1.79}
	// Para instância dessa forma, é preciso passar uma "pessoas" na criação do "estudante":
	e1 := estudante{p1, 123, "Análise e Desenvolvimento de Sistemas", "Bloco C"}
	// Considerando, a seguinte forma não funciona:
	// e2 := estudante{"Carlos", "Roberto", 28, 1.79, 123, "Análise e Desenvolvimento de Sistemas", "Bloco C"}

	e3 := estudante{} // Assim já é possível.
	e3.nome = "Carlinhos"
	e3.altura = 1.60

	e4 := estudante{pessoa: p1, curso: "Letras"}

	fmt.Println(p1)
	fmt.Println(e1)
	// fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(e4)
}
