package repositorios

import (
	"database/sql"
	"devbook/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
