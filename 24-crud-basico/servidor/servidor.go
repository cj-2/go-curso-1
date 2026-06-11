package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha na requisição."))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro na conversão."))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro na conexão com o banco de dados."))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar usuário."))
		return
	}
	defer statement.Close()

	fmt.Print(usuario)
}
