package auxiliar

import "fmt"

/*
O Go define a visibilidade de uma função ou variável pela primeira letra.

Escrever() => Pública
escrever() => Visível apenas no dentro do pacote.

*/

// Por padrão o Go espera que haja um comentário acima de uma função explicando sua função
func Escrever() {
	fmt.Println("Escrevendo da função auxiliar.")
	escrever()
}
