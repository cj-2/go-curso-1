package main

import "fmt"

// Não é preciso utilizar o := já que a variável já está definida no retorno.
func calculosMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(15, 10)
	fmt.Println(soma, subtracao)
}
