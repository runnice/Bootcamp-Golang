package main

import (
	"errors"
	"fmt"
	"os"
)

type cliente struct {
	arquivo   string
	nome      string
	sobrenome string
	rg        string
	telefone  string
	endereco  string
}

func (c *cliente) gerarID(id string) {
	if id == "" {
		panic("id não pode ser vazio")
	}
	c.arquivo = id
}

func (f *cliente) checarUsuario(arquivo string) error {
	defer debugger()
	arq, err := os.Open(arquivo)
	if err != nil {
		err = errors.New("o arquivo indicado não foi encontrado ou está danificado")
		panic(err)
	}
	fmt.Println("Arquivo lido com sucesso")
	defer arq.Close()
	return nil
}

func debugger() {
	err := recover()

	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	cliente := cliente{"", "Gugu", "Bacana", "123456789", "99999999", "Rua da Fazenda"}

	// fmt.Println(cliente)
	// cliente.gerarID("")
	// fmt.Println(cliente)

	cliente.checarUsuario("customers.txt")

}
