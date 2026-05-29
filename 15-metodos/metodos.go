package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

/*
	Declarando método que será acessível a partir
	de um struct usuario.
*/
func (u usuario) salvar() {
	fmt.Printf("Salvando... %s \n", u.nome)
}

/*
	Para alterar a struct é necessário passar por
	referencia utilizando o "*".
*/
func (u *usuario) atualizarNome(nome string) {
	u.nome = nome
}

func main() {
	user := usuario{"José", 18}
	user2 := usuario{"Carlos", 28}

	// Formas de chamar o método
	user.salvar()
	usuario.salvar(user2)

	user.atualizarNome("Juca")
	user.salvar()
}
