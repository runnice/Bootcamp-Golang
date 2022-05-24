package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type produto struct {
	Id         int
	Quantidade int
	Preco      float64
}

func (p produto) lineGeneratorCSV() string {
	return fmt.Sprintf("%d, %.2f, %d\n", p.Id, p.Preco, p.Quantidade)
}

func (p produto) lineHeaderCSV() string {
	return "ID, Preço, Quantidade\n"
}

func registraArquivo(nomeDoArquivo string, produtos []produto) error {
	arquivo, err := os.OpenFile(nomeDoArquivo, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo : %w", err)
	}
	defer arquivo.Close()
	p := produtos[0]

	if _, err = arquivo.WriteString(p.lineHeaderCSV()); err != nil {
		return fmt.Errorf("erro ao gerar primeira linhe: %w", err)
	}

	for _, produto := range produtos {
		if _, err = arquivo.WriteString(produto.lineGeneratorCSV()); err != nil {
			fmt.Println("error to save line: ", err)
		}
	}
	return nil

}

func lerArquivo(nomeDoArquivo string) {

	arquivo, err := os.Open(nomeDoArquivo)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
	}
	defer arquivo.Close()

	primeiraLinha := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	scanner := bufio.NewScanner(arquivo)

	scanner.Scan()

	header := strings.Split(scanner.Text(), ";")

	for _, c := range header {
		fmt.Fprintf(primeiraLinha, "%s\t\t", c)
	}
	fmt.Fprintln(primeiraLinha)

	var total float64
	for scanner.Scan() {
		valores := strings.Split(scanner.Text(), ",")
		for i, v := range valores {
			fmt.Fprintf(primeiraLinha, "%s\t\t", v)
			if i == 0 {
				preco, _ := strconv.ParseFloat(strings.TrimSpace(valores[1]), 64)
				quantidade, _ := strconv.Atoi(strings.TrimSpace(valores[2]))
				total += preco * float64(quantidade)
			}

		}

		fmt.Fprintln(primeiraLinha)
	}
	fmt.Fprintf(primeiraLinha, "\t\t%.2f\n", total)

	primeiraLinha.Flush()
}

func main() {

	var produtos = []produto{}
	produtos = append(produtos, produto{Id: 1, Quantidade: 10, Preco: 100})
	produtos = append(produtos, produto{Id: 2, Quantidade: 20, Preco: 200})
	produtos = append(produtos, produto{Id: 3, Quantidade: 30, Preco: 300})

	registraArquivo("produtos.csv", produtos)
	lerArquivo("produtos.csv")

}
