package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("address:[::]:8080: got / request\n")
	io.WriteString(w, "Aquecendo os motores!\n")
}

type Request struct {
	Base64 string `json:"base64"`
}

type Response struct {
	Msg string `json:"msg"`
}

func validateHeader(w http.ResponseWriter, r *http.Request) error {
	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return errors.New("Header inválido")
	}

	return nil
}

func upload(w http.ResponseWriter, r *http.Request) {
	err := validateHeader(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	bunmarshal, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req Request
	err = json.Unmarshal(bunmarshal, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res := Response{
		"Seu upload foi feito com sucesso!",
	}

	bmarshal, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	io.WriteString(w, string(bmarshal))
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/upload", upload)

	fmt.Println("======= Start server =======")
	fmt.Println("======= Listening on =======")

	_ = http.ListenAndServe(":8080", nil)
}
