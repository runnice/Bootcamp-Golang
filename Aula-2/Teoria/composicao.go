package main

// Composição - Conceito de herança
// Go não possui herança, mas sim composição. (embedding structs)
// Declaramos nossa estrutura base do Veículo e nela adicionaremos os campos de km e hora

import "fmt"

type Veiculo struct {
	km   float64
	hora float64
}

func (v Veiculo) detalhe() {
	fmt.Printf("%f km/hora\n", v.km/v.hora)
}

// Delcarar uma estrutura composta, a estrutura Automóvel. Nele vamos adicionar um campo do tipo Veículo.
type Automovel struct {
	v Veiculo
}

//Vamos adicionar um método que receba o tempo em minutos e cuide do cálculo da distâcia com base em 100km/h

func (a *Automovel) Correr(minutos int) {
	a.v.hora = float64(minutos) / 60
	a.v.km = 100
}

// E o método detalhe que chama o método da estrutura "pai"
func (a *Automovel) Detalhe() {
	fmt.Println("\nV: \tAutomóvel")
	a.v.detalhe()
}

// Declaramos nossa estrutura composta Moto

type Moto struct {
	v Veiculo
}

// Adicionamos o método Correr na estrutura Moto

func (m *Moto) Correr(minutos int) {
	m.v.hora = float64(minutos) / 60
	m.v.km = m.v.hora * 80
}

// E o método detalhe que chama o método da estrutura "pai"

func (m *Moto) Detalhe() {
	fmt.Println("\nV: \tMoto")
	m.v.detalhe()
}

func main() {
	// Finalmente executamos nossos métodos no main do projeto e vemos resultados:

	palio := Automovel{}
	palio.Correr(60)
	palio.Detalhe()

	kawasaki := Moto{}
	kawasaki.Correr(60)
	kawasaki.Detalhe()

}

//
