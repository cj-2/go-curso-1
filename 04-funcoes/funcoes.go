package main

import "fmt"

func somar(a int, b int) int8 {
	return int8(a + b)
}

// função com múltiplos retornos e declaração de tipo "resumida" nos parâmetros
func calculosMath(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(8, 15)
	fmt.Println(soma)

	// declarando função em uma variável
	var f = func() {
		fmt.Println("Plin!")
	}

	f()

	// _ define que não vou usar.
	resultado1, _ := calculosMath(10, 15)

	fmt.Println(resultado1)
}
