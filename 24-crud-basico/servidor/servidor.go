package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// Preprando o comando que será feito.
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement."))
		return
	}
	defer statement.Close()

	execResult, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement."))
		return
	}

	idUsuario, erro := execResult.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso. Id: %d", idUsuario)))
}

// BuscarUsuarios traz todos os usuarios do banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select id, nome, email from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao consultar no banco de dados"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário."))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao fazer encoder dos usuários."))
		return
	}

}

// BuscarUsuario traz um usuário do banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro."))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select id, nome, email from usuarios where id = ? limit 1", id)
	if erro != nil {
		w.Write([]byte("Erro ao consultar no banco de dados."))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário."))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao fazer encoder dos usuários."))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro."))
		return
	}

	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler body da requisição."))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro na conversão do body."))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement."))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, id); erro != nil {
		w.Write([]byte("Erro ao executar o statement."))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro."))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement."))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		w.Write([]byte("Erro ao executar o statement."))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
