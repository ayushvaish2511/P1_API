package config

import (
	"P1_API/helper"
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "postgres"
)

func DatabaseConnection() *sql.DB {
	// sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbName=%s sslmode=disable", host, port, user, password, dbName)

	db,err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable")
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database")

	return db

}