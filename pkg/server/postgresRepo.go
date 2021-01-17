package server

import (
	"context"

	"github.com/cepriego/test/pkg/postgres"
	"github.com/randallmlough/pgxscan"
)

type Postgresrepo struct {
	pgconn *postgres.PostgressConn
}

func NewPostgresRepo(pgcon *postgres.PostgressConn) *Postgresrepo {
	return &Postgresrepo{
		pgconn: pgcon,
	}
}

func (r *Postgresrepo) GetUsers() ([]Usuario, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return []Usuario{}, err
	}

	defer conn.Release()

	q := `SELECT "ID", "fechaCreacion", "Nombre"
	FROM public.usuarios
	where "Eliminado" != 'true'`

	rows, err := conn.Query(context, q)
	if err != nil {
		return []Usuario{}, err
	}

	defer rows.Close()

	users := []Usuario{}
	err = pgxscan.NewScanner(rows).Scan(&users)

	if err != nil {
		return []Usuario{}, err
	}
	return users, nil
}

func (r *Postgresrepo) GetUser(id int) ([]Usuario, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return []Usuario{}, err
	}

	defer conn.Release()

	q := `SELECT "ID", "fechaCreacion", "Nombre", "Eliminado"
	FROM public.usuarios
	where "ID" = $1`

	rows, err := conn.Query(context, q, id)
	if err != nil {
		return []Usuario{}, err
	}

	defer rows.Close()

	users := []Usuario{}
	err = pgxscan.NewScanner(rows).Scan(&users)

	if err != nil {
		return []Usuario{}, err
	}
	return users, nil
}

func (r *Postgresrepo) CreateUser(user Usuario) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()
	q := `INSERT INTO public.usuarios(
	"ID", "fechaCreacion", "Nombre", "Eliminado")
	VALUES ($1, $2, $3, $4);`

	rows, err := conn.Query(context, q, user.Id, user.FechaCreacion, user.Nombre, user.Eliminado)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}

func (r *Postgresrepo) CreateGroup(grupo Grupo) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()
	q := `INSERT INTO public."Grupos"(
		"ID", "FechaCreacion", "Nombre", "Eliminado")
		VALUES ($1, $2, $3, $4);`

	rows, err := conn.Query(context, q, grupo.Id, grupo.FechaCreacion, grupo.Nombre, grupo.Eliminado)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}

func (r *Postgresrepo) GetGroup(id int) ([]Grupo, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return []Grupo{}, err
	}

	defer conn.Release()

	q := `SELECT "ID", "FechaCreacion", "Nombre", "Eliminado"
	FROM public."Grupos"
	where "ID" = $1`

	rows, err := conn.Query(context, q, id)
	if err != nil {
		return []Grupo{}, err
	}

	defer rows.Close()

	grupos := []Grupo{}
	err = pgxscan.NewScanner(rows).Scan(&grupos)

	if err != nil {
		return []Grupo{}, err
	}
	return grupos, nil
}

func (r *Postgresrepo) AssignUserToGroup(usrGrp UsuarioGrupo) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()

	q := `INSERT INTO public."Grupos_Usuarios"(
		id_usuario, id_grupo)
		VALUES ($1, $2);`

	rows, err := conn.Query(context, q, usrGrp.IdUsuario, usrGrp.IdGrupo)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}

func (r *Postgresrepo) GetGroupsAndUsers() ([]GroupWithUsers, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return []GroupWithUsers{}, err
	}

	defer conn.Release()

	q := `SELECT gr."ID" as "ID", gr."Nombre" as "Grupo", usr."Nombre" as "Usuario"
	FROM public."Grupos" gr
	INNER JOIN public."Grupos_Usuarios" grpus on gr."ID" = grpus."id_grupo"
	INNER JOIN usuarios usr on  usr."ID" = grpus."id_usuario"
	WHERE gr."Eliminado" = 'false' AND usr."Eliminado" = 'false'`

	rows, err := conn.Query(context, q)
	if err != nil {
		return []GroupWithUsers{}, err
	}

	defer rows.Close()

	grupos := []GroupWithUsers{}
	err = pgxscan.NewScanner(rows).Scan(&grupos)

	if err != nil {
		return []GroupWithUsers{}, err
	}
	return grupos, nil
}

///Hides a user, the user is not deleted from the db, users are only marked as deleted
func (r *Postgresrepo) DeleteUser(id int) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()
	q := `UPDATE public.usuarios
	SET "Eliminado" = 'true'
	WHERE "ID" = $1;`

	rows, err := conn.Query(context, q, id)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}

//Deletes the relationship of user and group from the table.
func (r *Postgresrepo) DeleteUserFromGroup(idUser int, idGroup int) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()
	q := `DELETE FROM public."Grupos_Usuarios"
	WHERE "id_usuario" = $1 AND "id_grupo" = $2`

	rows, err := conn.Query(context, q, idUser, idGroup)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}

//Deletes the relationship of user and group from the table.
func (r *Postgresrepo) DeleteGroup(idGroup int) (bool, error) {
	context := context.Background()
	conn, err := r.pgconn.GetConn(context)
	if err != nil {
		return false, err
	}

	defer conn.Release()
	q := `UPDATE public."Grupos"
	SET "Eliminado" = 'true'
	WHERE "ID" = $1;`

	rows, err := conn.Query(context, q, idGroup)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil
}
