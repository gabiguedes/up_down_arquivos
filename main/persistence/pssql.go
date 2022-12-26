package persistence

import "fmt"

type Arquivo struct {
	ID    int
	Corpo []byte
}

const (
	PostgresDriver = "postgres"
	host           = "localhost"
	port           = "5432"
	user           = "postgres"
	password       = "123"
	Dbname         = "files"
)

var Datasource = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, Dbname)
