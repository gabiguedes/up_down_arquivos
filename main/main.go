package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("address:[::]:8080: got / request\n")
	io.WriteString(w, "Aquecendo os motores!\n")
}

type Request struct {
	Arquivo []byte `json:"arquivo" validate:"required"`
}

type Response struct {
	Msg string `json:"msg"`
}

func upload(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	res := Response{
		"deu bom demais",
	}

	fmt.Println("response ", res.Msg)
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("marshal ", string(b))

	io.WriteString(w, string(b))
	fmt.Printf("body: %s", body)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/arquivo", upload)

	fmt.Println("======= Start server =======")
	fmt.Println("======= Listening on =======")

	_ = http.ListenAndServe(":8080", nil)
}
