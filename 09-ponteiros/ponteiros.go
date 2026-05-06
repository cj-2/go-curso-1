package main

import "fmt"

func main() {

	var var1 int = 10
	var var2 int = var1 // copiando valor, não referência.

	fmt.Println(var1, var2) // 10 10
	var1++
	fmt.Println(var1, var2) // 11 10

	var var3 int = 100
	var ponteiro *int
	fmt.Println(var3, ponteiro) // 100 <nil>

	ponteiro = &var3
	fmt.Println(var3, ponteiro)  // 100 0x151b113e60a8 (endereço em memória)
	fmt.Println(var3, *ponteiro) // 100 100
	var3++
	fmt.Println(var3, *ponteiro) // 101 101
}
