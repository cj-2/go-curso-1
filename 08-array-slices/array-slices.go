package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string // Obrigatório o tamanho, ou se tornará um slice.
	array1[0] = "Bolinha"
	array2 := [5]string{"X"} // Declarando por inferência.

	// fixando o tamanho dinâmicamente baseado na quantidade de itens.
	array3 := [...]int{1, 2, 3, 4, 5} // O [...] que faz isso.
	// porem, não deixa "dinâmico" depois, apenas na inicialização!

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	// slices são muito utilizados no Go
	slice1 := []int{1, 2, 3} // slices são definidos assim, sem quantidade de "espaços"
	fmt.Println(slice1)

	// Operações slice
	slice1 = append(slice1, 4) // retorna um "novo slice" para a referência inicial.
	fmt.Println(slice1)

	// fatiando arrays, o nome referência um "pedaço" do array:
	slice2 := array3[0:3] // indice 1 inclusivo, indice 2 exclusivo (não é incluído)
	fmt.Println(slice2)   // [1 2 3]

	array3[2] = 200
	fmt.Println(slice2) // [1 2 200]

	fmt.Println(reflect.TypeOf(array1)) // [5]string
	fmt.Println(reflect.TypeOf(slice1)) // []int
}
