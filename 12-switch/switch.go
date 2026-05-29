package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Dia inválido"
	}
}

/*
	Não é passado para o switch.
*/
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sabado"
	default:
		return "Dia inválido"
	}
}

func fallSwitch(numero int) string {
	var dia string
	switch {
	case numero == 1:
		fmt.Println("Só deveria cair aqui... Mas utilizei um fallthrough!")
		dia = "Domingo"
		fallthrough // vai jogar o código para a próxima condição
	case numero == 2:
		fmt.Println("Pim salabim, rodei mesmo assim!")
	default:
		fmt.Println("Nenhuma emoção aqui.")
	}

	return dia
}

func main() {
	dia := diaDaSemana(3)

	fmt.Println(dia)
	fmt.Println(diaDaSemana(4))

	fallSwitch(1)
}
