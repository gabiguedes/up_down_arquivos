package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter) {
	fmt.Printf("address:[::]:8080: got / request\n")
	io.WriteString(w, "warming up the engines!\n")
}

type Request struct {
	Base64 string `json:"base64" json:"required" json:"base64"`
}

type Response struct {
	Msg string `json:"msg"`
}

func ValidateHeader(r *http.Request) error {
	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		return errors.New("header required")
	}

	return nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) *Request {
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

func HandlerResponse(w http.ResponseWriter) map[string]string {
	response := map[string]string{"msg": "Your upload was successful!"}
	_, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		io.WriteString(w, err.Error())
		return nil
	}

	return response
}
