package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Struct que vai receber o banco, a conexão vai ser aberta no package controllers e passada no pacote repositorios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}