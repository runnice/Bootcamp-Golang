package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Product struct {
	Id         int
	Quantidade int
	Preco      float64
}

func registraProdutoComprado(p Product) {
	arquivo, err := os.OpenFile("produtosComprados.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()
	str := fmt.Sprintf("%d;%d;%.2f;\n", p.Id, p.Quantidade, p.Preco)
	arquivo.WriteString(str)
}

func imprimeArquivo(nomeDoArquivo string) {
	arquivo, err := ioutil.ReadFile(nomeDoArquivo)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

// func imprimeArquivo2(nomeDoArquivo string) ([]Product, error) {
// 	f, err := os.Open(nomeDoArquivo)
// 	if err != nil {
// 		return []Product{}, err
// 	}
// 	defer f.Close()
// 	rows := csv.NewReader(f)
// 	var ps []Product
// 	for {
// 		row, err := rows.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return ps, err
// 		}

// 	}
// 	return ps, nil
// }

func main() {

	camisa := Product{1, 10, 29.90}
	bola := Product{2, 4, 19.90}

	registraProdutoComprado(camisa)
	registraProdutoComprado(bola)

	imprimeArquivo("produtosComprados.csv")
}
