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

	req := server.HandleRequest(w, r)
	psInsert(req.Base64)
	res := server.HandlerResponse(w)

	io.WriteString(w, res["msg"])
}

func download(w http.ResponseWriter, r *http.Request) {
	res := psSelectPorID()

	io.WriteString(w, string(res.Corpo))
}

func psInsert(base64 string) {
	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2)", "arquivo")

	insert, err := db.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

	result, err := insert.Exec(2, base64)
	if err != nil {
		fmt.Println(err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(affect)
}

func psSelectPorID() *persistence.Arquivo {
	var arquivo persistence.Arquivo

	sqlStatement := fmt.Sprintf("SELECT id, corpo FROM %s where id = $1", "arquivo")

	err := db.QueryRow(sqlStatement, 2).Scan(&arquivo.ID, &arquivo.Corpo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d\t%s \n", arquivo.ID, arquivo.Corpo)

	return &arquivo
}

func psSelect() {
	state, err := db.Query("select id, corpo from arquivo")
	if err != nil {
		panic(err)
	}

	for state.Next() {
		var arquivo persistence.Arquivo
		err = state.Scan(&arquivo.ID, &arquivo.Corpo)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", download)
	fmt.Println("======= Start server =======")

	fmt.Printf("Accessing [%s] database... ", persistence.Dbname)
	db, _ = sql.Open(persistence.PostgresDriver, persistence.Datasource)

	defer db.Close()

	psSelect()

	_ = http.ListenAndServe(":8080", nil)
}
