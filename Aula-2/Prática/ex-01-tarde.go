package main

import "fmt"

type Aluno struct {
	Nome         string
	Sobrenome    string
	RG           int
	DataAdmissao string
}

func (a Aluno) ImprimeAlunos() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("Data de admissão:", a.DataAdmissao)
}

func main() {

	AlunosBootcamp := []Aluno{
		{"João", "da Silva", 123456789, "01/01/2019"},
		{"Maria", "da Silva", 123456789, "01/01/2019"},
		{"José", "da Silva", 123456789, "01/01/2019"},
		{"Pedro", "da Silva", 123456789, "01/01/2019"},
	}

	for _, aluno := range AlunosBootcamp {
		aluno.ImprimeAlunos()
	}
}
