package main

import "fmt"

const (
	pequeno                = "pequeno"
	medio                  = "medio"
	grande                 = "grande"
	custoPequenoEmPorcento = 0
	custoMedioEmPorcento   = 3
	custoGrandeEmPorcento  = 6
	custoEnvioGrande       = 2500
)

type loja struct {
	produtos []Produto
}

type produto struct {
	Nome          string
	Preco         float64
	TipoDeProduto string
}

type Produto interface {
	CalculaCusto() float64
}

type Ecommerce interface {
	AdicionaProduto(p Produto)
	Total() float64
}

func novoProduto(nome string, preco float64, tipo string) Produto {
	novoProduto := produto{nome, preco, tipo}
	return &novoProduto

}

func novaLoja() Ecommerce {
	var novaLoja = loja{}
	return &novaLoja
}

func (p *produto) CalculaCusto() float64 {
	switch p.TipoDeProduto {
	case pequeno:
		return p.Preco * (float64(custoPequenoEmPorcento)/100 + 1)
	case medio:
		return p.Preco * (float64(custoMedioEmPorcento)/100 + 1)
	case grande:
		return p.Preco*(float64(custoGrandeEmPorcento/100+1)) + custoEnvioGrande
	default:
		return p.Preco
	}
}

func (l *loja) AdicionaProduto(p Produto) {
	l.produtos = append(l.produtos, p)

}

func (l *loja) Total() (precoTotal float64) {
	for _, produto := range l.produtos {
		precoTotal += produto.CalculaCusto()
	}
	return precoTotal
}

func main() {

	maca := novoProduto("Maca", 5, pequeno)
	melancia := novoProduto("Melancia", 10, medio)
	bananaDeOuro := novoProduto("Banana de ouro", 200, grande)

	fmt.Println("Preço da maçã é: R$", maca.CalculaCusto())
	fmt.Println("Preço da melancia é: R$", melancia.CalculaCusto())
	fmt.Println("Preço da Banana De Ouro é: R$", bananaDeOuro.CalculaCusto())

	quitandaDoJoao := novaLoja()
	quitandaDoJoao.AdicionaProduto(maca)
	quitandaDoJoao.AdicionaProduto(melancia)
	quitandaDoJoao.AdicionaProduto(bananaDeOuro)

	fmt.Println(quitandaDoJoao.Total())

}
