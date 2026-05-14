package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Carlos Roberto",
		"sobrenome": "da Silva Júnior",
	}

	// map[nome:Carlos Roberto sobrenome:da Silva Júnior]
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// map alinhado:
	usuarios := map[string]map[string]string{
		"carlim88": {
			"nome":  "Carlos",
			"idade": "28",
		},
		"jojo": {
			"nome":  "José",
			"idade": "36",
		},
	}

	// map[carlim88:map[idade:28 nome:Carlos] jojo:map[idade:36 nome:José]]
	fmt.Println(usuarios)
	delete(usuarios, "jojo") // apagando item por chave
	//map[carlim88:map[idade:28 nome:Carlos]]
	fmt.Println(usuarios)

	// adicionando item
	usuarios["mac"] = map[string]string{"nome": "Marcos"}
	fmt.Println(usuarios)
}
