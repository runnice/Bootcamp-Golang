package main

import "fmt"

type produtos struct {
	nome  string
	preco float64
	qntd  int
}

type servicos struct {
	nome               string
	preco              float64
	minutosTrabalhados int
}

type manutencao struct {
	nome  string
	preco float64
}

func somarProdutos(p []produtos, out chan<- float64) {
	total := 0.0
	for _, value := range p {
		total = value.preco * float64(value.qntd)
	}
	out <- total
	close(out)
}
func somarServicos(s []servicos, out chan<- float64) {
	total := 0.0
	for _, value := range s {
		if value.minutosTrabalhados < 30 {
			value.minutosTrabalhados = 30
			total += value.preco * float64(value.minutosTrabalhados)
		} else {
			total += value.preco * float64(value.minutosTrabalhados)
		}
	}
	out <- total
	close(out)
}

func somarManutencao(m []manutencao, out chan<- float64) {
	total := 0.0
	for _, value := range m {
		total += value.preco
	}
	out <- total
}

func main() {

	var p = []produtos{}

	p = append(p, produtos{nome: "Vela", preco: 5.0, qntd: 5})
	p = append(p, produtos{"Pneu", 100, 5})
	p = append(p, produtos{"Pneu", 100, 5})
	p = append(p, produtos{"Pneu", 100, 5})
	p = append(p, produtos{"Pneu", 100, 5})
	p = append(p, produtos{"Pneu", 100, 5})

	var s = []servicos{}
	s = append(s, servicos{"Troca de Vela", 0, 5})
	s = append(s, servicos{"Troca de Pneu", 100, 60})

	var m = []manutencao{}
	m = append(m, manutencao{"Troca de Vela", 100})
	m = append(m, manutencao{"Troca de Pneu", 100})

	canalSP := make(chan float64)
	canalSS := make(chan float64)
	canalSM := make(chan float64)

	go somarProdutos(p, canalSP)
	go somarServicos(s, canalSS)
	go somarManutencao(m, canalSM)

	totalSP := <-canalSP
	totalSS := <-canalSS
	totalSM := <-canalSM

	fmt.Println("Total de SomaProdutos", totalSP, totalSS, totalSM)
	fmt.Println("Total de SomaServiços", totalSS)
	fmt.Println("Total de SomaTotal", totalSM)

	fmt.Println("Total geral", totalSP+totalSS+totalSM)
}
