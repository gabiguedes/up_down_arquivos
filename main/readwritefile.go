package main

import (
	"fmt"
	"os"
)

const (
	ARQUIVO_LEITURA = "C:/Users/Gabi Guedes/Documents/uploadarquivos/leitura.txt"
	ARQUIVO_ESCRITA = "C:/Users/Gabi Guedes/Documents/uploadarquivos/escrita.txt"
)

func readFile(nomearquivo string) ([]byte, error) {
	rd, err := os.ReadFile(nomearquivo)
	if err != nil {
		fmt.Printf("Ocorre um erro na leitura do arquivo - Details %s", err.Error())
		panic(err)
	}
	return rd, nil
}

func writeFile(data []byte) error {
	err := os.WriteFile(ARQUIVO_ESCRITA, data, 0644)
	if err != nil {
		fmt.Printf("algo deu ruim... %s", err.Error())
		panic(err)
	}
	return nil
}

func together() {
	rd, _ := readFile(ARQUIVO_LEITURA)

	data := []byte("Loucura é querer resultados diferentes fazendo tudo exatamente igual")
	_ = writeFile(data)

	fmt.Printf("O arquivo %s contém a seguinte mensagem:\n %s", ARQUIVO_LEITURA, rd)
	fmt.Printf("O arquivo %s contém a seguinte mensagem:\n %s", ARQUIVO_ESCRITA, rd)
}
