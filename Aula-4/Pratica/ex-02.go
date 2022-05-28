package main

import "fmt"

type user struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

type produto struct {
	nome  string
	preco float64
	qntd  int
}

func novoProduto(nome string, preco float64) produto {
	return produto{
		nome:  nome,
		preco: preco,
		qntd:  0,
	}
}

func adicionarProduto(u *user, p *produto, qntd int) {
	p.qntd = qntd
	u.produtos = append(u.produtos, *p)
}

func deletarProdutos(u *user) {
	u.produtos = []produto{}
}

func main() {
	u := user{
		nome:      "João",
		sobrenome: "Silva",
		email:     "blablabla@.gmail.com",
	}
	// fmt.Println(u)
	p := novoProduto("Coca-Cola", 5.0)
	// fmt.Println(u, p)

	adicionarProduto(&u, &p, 5)
	fmt.Println(u)

	deletarProdutos(&u)
	fmt.Println(&u)

}
