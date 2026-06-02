package main

import "fmt"

func main() {
	canal := make(chan string, 1)
	canal <- "Olá"

	mensagem := <-canal
	fmt.Println(mensagem)
}
