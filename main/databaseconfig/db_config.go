package databaseconfig

import "fmt"

type Files struct {
	file string
}

const (
	PostgresDriver = "postgres"
	host           = "localhost"
	port           = "5432"
	user           = "postgres"
	password       = ""
	Dbname         = "files"
)

var Datasource = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, Dbname)
