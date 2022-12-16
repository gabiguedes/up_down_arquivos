package main

import (
	"fmt"
	"os"
)

const (
	NOME_ARQUIVO_LER = "C:/Users/Gabi Guedes/Documents/uploadarquivos/ler.txt"
	CRIA_ARQUIVO     = "C:/Users/Gabi Guedes/Documents/uploadarquivos/test.txt"
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
	err := os.WriteFile(CRIA_ARQUIVO, data, 0644)
	if err != nil {
		fmt.Printf("algo deu ruim... %s", err.Error())
		panic(err)
	}
	return nil
}

func main() {
	rd, _ := readFile(NOME_ARQUIVO_LER)

	data := []byte("Loucura é querer resultados diferentes fazendo tudo exatamente igual")
	_ = writeFile(data)

	fmt.Printf("O arquivo %s contém a seguinte mensagem:\n %s", NOME_ARQUIVO_LER, rd)
	fmt.Printf("O arquivo %s contém a seguinte mensagem:\n %s", CRIA_ARQUIVO, rd)
}
