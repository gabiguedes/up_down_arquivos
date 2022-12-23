package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"uploadarquivos/main/databaseconfig"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("address:[::]:8080: got / request\n")
	io.WriteString(w, "warming up the engines!\n")
}

type Request struct {
	Base64 string `json:"base64" json:"required" json:"base64"`
}

type Response struct {
	Msg string `json:"msg"`
}

func validateHeader(r *http.Request) error {
	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		return errors.New("header required")
	}

	return nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) *Request {
	unmarshal, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		io.WriteString(w, err.Error())
		return nil
	}

	var req Request
	m := make(map[string]string)

	err = json.Unmarshal(unmarshal, &m)
	if err != nil {
		http.Error(w, err.Error(), 422)
		io.WriteString(w, err.Error())
		return nil
	}

	req.Base64 = m["base64"]
	return &req
}

func handlerResponse(w http.ResponseWriter) map[string]string {
	response := map[string]string{"msg": "Your upload was successful!"}
	_, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		io.WriteString(w, err.Error())
		return nil
	}

	return response
}

func upload(w http.ResponseWriter, r *http.Request) {
	err := validateHeader(r)
	if err != nil {
		http.Error(w, "Content-Type application/json is required", http.StatusUnsupportedMediaType)
		io.WriteString(w, err.Error())
		return
	}

	handleRequest(w, r)
	res := handlerResponse(w)

	io.WriteString(w, res["msg"])
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/upload", upload)

	fmt.Println("======= Start server =======")

	fmt.Printf("Accessing [%s] database... ", databaseconfig.Dbname)

	db, err := sql.Open(databaseconfig.PostgresDriver, databaseconfig.Datasource)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	_ = http.ListenAndServe(":8080", nil)
}
