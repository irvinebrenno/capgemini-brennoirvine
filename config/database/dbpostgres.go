package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// definição das constantes de acesso ao banco
const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "capgemini_brennoirvine"
)

// cria e retorna uma conexão com o bando de dados postgres
func Conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
