package model

import "fmt"

type Usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func NovoUsuario(usuario *Usuario) error {
	comando := fmt.Sprintf("insert into usuarios(nome, email, senha) values('%s', '%s', '%s');", usuario.Nome, usuario.Email, usuario.Senha)
	_, err := db.Exec(comando)
	return err
}

func ObterUsuario(id string) (Usuario, error) {
	var usuario Usuario
	comando := fmt.Sprintf("select * from usuarios where id = %s;", id)
	linhas, err := db.Query(comando)
	if err != nil {
		return Usuario{}, err
	}

	for linhas.Next() {
		err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)
		if err != nil {
			return Usuario{}, err
		}
	}
	return usuario, nil
}

func VerificaEmail(email string, usuario *Usuario) bool {
	comando := fmt.Sprintf("select id, nome, email, senha from usuarios where email = '%s'", email)
	fmt.Println(comando)
	rows, err := db.Query(comando)
	if err != nil {
		fmt.Println("erro na execução", err)
		return false
	}
	fmt.Println("linhas", rows)

	for rows.Next() {
		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)
		if err != nil {
			return false
		} else {
			return true
		}
	}
	fmt.Println("Chegou ao fim", err)
	return false
}
