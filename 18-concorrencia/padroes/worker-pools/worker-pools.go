package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

/*
Faz a iteração das tarefas do canal (calculo de um número no fibonacci)
E repassa para o canal "resultados" o que foi calculado.
*/
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Chamamos o worker, mesmo sem nada por enquanto.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Populando o canal de tarefas com os números a serem calculados.
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	// Fechamos então o canal de tarefas.
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	close(resultados)
}
