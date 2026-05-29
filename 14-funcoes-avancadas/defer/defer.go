package main

import "fmt"

func fun1() {
	fmt.Println("Função 1")
}

func fun2() {
	fmt.Println("Função 2")
}

func main() {
	defer fun1()
	fun2()
}
