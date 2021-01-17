package server

import (
	"time"
)

type Usuario struct {
	Id            int       `db:"ID"`
	FechaCreacion time.Time `db:"fechaCreacion"`
	Nombre        string    `db:"Nombre"`
	Eliminado     bool      `db:"Eliminado"`
}

type Grupo struct {
	Id            int       `db:"ID"`
	FechaCreacion time.Time `db:"FechaCreacion"`
	Nombre        string    `db:"Nombre"`
	Eliminado     bool      `db:"Eliminado"`
}

type UsuarioGrupo struct {
	IdUsuario int `db:"id_usuario"`
	IdGrupo   int `db:"id_grupo"`
}

type GroupWithUsers struct {
	Id      int    `db:"ID"`
	Grupo   string `db:"Grupo"`
	Usuario string `db:"Usuario"`
}
