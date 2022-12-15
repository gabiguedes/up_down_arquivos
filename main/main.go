package main

import (
	"fmt"
	"os"
)

const NOME_ARQUIVO = "C:/Users/Gabi Guedes/Documents/uploadarquivos/ler.txt"

func readFile(nomearquivo string) ([]byte, error) {
	rd, err := os.ReadFile(nomearquivo)
	if err != nil {
		fmt.Printf("Ocorre um erro na leitura do arquivo - Details %s", err.Error())
		panic(err)
	}
	return rd, nil
}

func main() {
	rd, _ := readFile(NOME_ARQUIVO)

	fmt.Printf("O arquivo contém a seguinte mensagem:\n %s", rd)

}
