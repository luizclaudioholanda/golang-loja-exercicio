package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDeDados() *sql.DB {
	conexao := "user=cdelivery_dev dbname=postgres password=cdelivery_dev host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
