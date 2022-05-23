package main

// Aula tarde - Go Bases

// Estrutura de dados - Rótulos e Tags - Interface - Metodos dentro de estrutura (Função receiver)

// Estrutura é uma coleção de dados. Ex: Pessoa, nela tem nome, idade, endereço, etc...

// Estrutura de Pessoa
// PS: Com as variávies em com primeira letra em Maíusculo, o atributo fica disponível para todos os métodos.

import (
	"encoding/json"
	"fmt"
	"reflect"

	circ "github.com/runnice/Bootcamp-Golang/Aula-2/Aula-2/Teoria/pkg/circulo"
)

// Definição de estrutura "Struct"
// para instanciarmos uma estrutura precisamos atribuir a uma variável
type Pessoa struct {
	Nome      string
	Genero    string
	Idade     int
	Profissao string
	Peso      float64
	Gostos    Preferencias //Colocar uma Struct dentro de uma struct
}

// vinicius:= Pessoa{"'Vinicius', 'Masculino', 35, "Software Developer", 90.5} -
// Podemos realiar uma aribuição direta ex: vinicius.Nome = "Vinicius"
// type Pessoas []Pessoa{} cria um Array de pessoas

// Definimios uma nova estrutura e colocamos ela dentro da Estrutura de Pessoa
type Preferencias struct {
	Comidas  string
	Filme    string
	Series   string
	Animes   string
	Esportes string
}

// Para acessarmos e modificarmos essa estrutura fazemos:
// vinicius.Gostos.Comidas = "Frango"
// ou podemos adicionar diretamente a estrutura:
// vicente := Pessoa{
// 	nome: "Vicente"
// 	genero: "Masculino"
// 	idade: 35
// 	profissao: "Software Developer"
// 	peso: 90.5
// 	gostos: Preferencias{Comidas: "Frango", Filme: "Titanic", Series: "Friends", Animes: "DBZ", Esportes: "Futebol"}
// }

// Exemplo de uma função que recebe uma estutura como parametro e exibe os dados dela:
func exibirDados(pessoa Pessoa) {
	fmt.Println(pessoa)

}

//Rotulos ou tags - Desta forma estamos dizendo se formos usar json, vai aparecer conforme rotulamos.
//bson para o mongodb

type Pessoa2 struct {
	PrimeiroNome string `bd:"primeiro_nome]"`
	SegundoNome  string `bd:"segundoNome"`
	Telefone     int    `bd:"telefone"`
	Endereco     string `bd:"endereco"`
}

// Pacote Reflect obter info da nossa estrutura em tempo de execução

// p:=reflect.TypeOf(pessoa)

//Exemplo de conversão de JSON na linha 143

//Metodos

// type Circulo struct {
// 	raio float64
// }

// func (c Circulo) SetarRaio(raio float64) { // Função Receiver é um parametro antes da função que recebe uma estrutura
// 	c.raio = raio
// 	fmt.Println("Raio:", c.raio)
// }
// Herença - Em Go não existe herança, mas sim composição.

// Ponteiro - Quando criamos uma função e passamos uma variável como argumento, o que a função faz
// é copiar o valor da variável e trabalhar com esse valor. Portanto, a variável que passamos como argumento
// NÃO é modificada.

// func (c *Circulo) setRaio(raio float64) { //
// 	c.raio = raio

// dessa forma c.raio será alterado.

// Composição - Conceito de herança
// Go não possui herança, mas sim composição. (embedding structs)
// Declaramos nossa estrutura base do Veículo e nela adicionaremos os campos de km e hora
type Veiculo struct {
	km   float64
	hora float64
}

func (v Veiculo) detalhe() {
	fmt.Printf("%f km/hora\n", v.km/v.hora)
}

type Automovel struct {
	v Veiculo
}

//Interfaces
// Uma Interface é uma forma de defifnir um contrato entre duas ou mais estruturas.

func main2() {
	vinicius := Pessoa{
		Nome:      "Vinicius",
		Genero:    "Masculino",
		Idade:     25,
		Profissao: "Programador",
		Peso:      80.5,
	}
	exibirDados(vinicius)

	joaquim := Pessoa{} // Cria uma estrutura vazia

	// atribui valores aos atributos após ter declarado vazio
	joaquim.Nome = "Joaquim"
	joaquim.Genero = "Masculino"

	// associar uma struct a uma variável sem declaração curta
	var maria Pessoa
	maria.Nome = "Maria"
	maria.Genero = "Feminino"

	// Duas formas de atribuir a Structs
	p1 := Pessoa{"Vinicius", "Masculino", 25, "Programador", 80.5, Preferencias{"frango", "titanic", "", "", ""}}
	p2 := Pessoa{Nome: "Joana", Genero: "Feminino", Idade: 20, Profissao: "Estudante", Peso: 50.5, Gostos: Preferencias{"Frango", "Lord of the Rings", "Friends", "DBZ", "Ping Pong"}}

	// Pode usar _ = Pessoa {} para ignorar variavel e n usar

	exibirDados(p1)
	exibirDados(p2)

	ppl := Pessoa2{"Paula", "Monteiro", 434343, "Rua do Limoeiro 321"}

	// Exemplo de Conversão de JSON com rótulos.
	meuJSON, err := json.Marshal(ppl)
	if err != nil {
		fmt.Println("Erro ao gerar JSON")
	}
	fmt.Println(string(meuJSON))

	people := Pessoa2{}
	p := reflect.TypeOf(people)

	fmt.Println("Type", p.Name())
	fmt.Println("Type", p.Kind())

	//Com o método NumField Podemos obter o número de campos que temos em nossa estrutura, isso nos ajudará com a estrutura
	for i := 0; i < p.NumField(); i++ {
		field := p.Field(i)
		tag := field.Tag.Get("bd")
		fmt.Println(tag)
	}

	//Métodos
	// Go não possui classes. Temos que instanciar a estrutura para usar os métodos
	circulo := circ.Circulo{} // instanciou a struct (como um objeto no java)
	circulo.SetarRaio(5)
	fmt.Println(circulo.GetRaio())

	//Ponteiros

	// Composição - Conceito de herança
	// Go não possui herança, mas sim composição. (embedding structs)

}
