package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"uploadarquivos/main/persistence"
	"uploadarquivos/main/server"
)

var db *sql.DB

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("address:[::]:8080: got / request\n")
	io.WriteString(w, "warming up the engines!\n")
}

func upload(w http.ResponseWriter, r *http.Request) {
	err := server.ValidateHeader(r)
	if err != nil {
		http.Error(w, "Content-Type application/json is required", http.StatusUnsupportedMediaType)
		io.WriteString(w, err.Error())
		return
	}

	server.HandleRequest(w, r)
	res := server.HandlerResponse(w)

	io.WriteString(w, res["msg"])
}

func psSelect() {
	state, err := db.Query("select id, corpo from arquivo")
	if err != nil {
		panic(err)
	}

	for state.Next() {
		var arquivo persistence.Arquivo
		err = state.Scan(&arquivo.ID, &arquivo.Corpo)

		fmt.Printf("\n RESULTS [%d\t%s] \n", arquivo.ID, arquivo.Corpo)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/upload", upload)
	fmt.Println("======= Start server =======")

	fmt.Printf("Accessing [%s] database... ", persistence.Dbname)
	db, _ = sql.Open(persistence.PostgresDriver, persistence.Datasource)

	defer db.Close()

	psSelect()

	_ = http.ListenAndServe(":8080", nil)
}
