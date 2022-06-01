package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// O que é uma API

// API é um conjunto de definicções e protocolos usados para desenvolver e integrar softwares
// de aplicativo. API significa Interface de Programação de Aplicativos.

// Eles não são  a parte visível, mas o circuito interno que
// apenas os desenvolvedores veem e conectam para fazer uma ferramenta funcionar

// SOPA protcolo XML

// Uma API pode ser pública ou privada (privada é a mais segura)
// AS APIS permitem que seus produtos e serviços se comuniquem com outros, sem a ecessidade
// de saber como elas são implementadas. Elas concedem flexibilidade; eles simplificam o design,
// o gerenciamento e uso das aplicações

//HTTP - Protocolo comunicação cleinte e servidor

// METODOS HTTP ( GET - POST - PUT - DELETE )

// ARQUITETURA WEB - possuem 6 categorias

// Cliente-servidor, Interface uniforme, Sistema de camadas, Cache, Statless e Código sob Demanada
// end-point : Ponto final de acesso
// ex: globo.com/contato - contato é o end-point

// REST - REpresentational State Transfer
// Padrão de desenvolvimento de API usando JSON como forma de comunicação
// JSON é um formato de troca de dados conveniente devido a sua facilidade de leitura

// Sintaxe JSON - coleção de pares nome/valor. Uma lista ordenada de valores.

// JSON aceita: Números, string, booleanos, nulos, listas, objetos e arrays.

// Marshal() - Empacotar estrutura GO e transformar em JSON

// omityempty retorna vazio, caso esteja vazio.

// Exemplo de Marshal sendo chamado do 55 ao 64

// Unmarshal() - Caso ele não consiga fazer o desempacotamento, ele retorna um erro.

// PACKAGE net/HTTP
// Pacote permite manipulações de rotas com servidor.

type product struct {
	Name      string `json:"name"`
	Price     int    `json:"price",omityempty:"true"`
	Published bool   `json:"published"`
	// Exemplo de como pegar quando não sabe o tipo

	// User struct{
	// 	Name any
	// 	CPF json.RawMessage

	// 	`json:"nome"`
	// }
}

// Gin - MicroFramework de desenvolvimento de APIs
// Como funciona o GIN
// Request -> Router -> Middleware -> Route Handler -> Middleware -> Response

// objtendo Gin go get -u github.com/gin-gonic/gin
// usando o Gin import "github.com/gin-gonic/gin"

func main() {
	// mapa pra pegar tudo sem mapear
	//c:= make(map[string]any)

	p := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}
	jsonx, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonx))

	// Exemplo de Unmarshal sendo chamado do 65 ao 72

	jsonData := `{"Name":"MacBook Pro", "Price":1500, "Published":true}`

	var p2 product

	if err2 := json.Unmarshal([]byte(jsonData), &p2); err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(p2)
	fmt.Println("Name ", p2.Name)
	fmt.Println("Price ", p2.Price)
	fmt.Println("Published ", p2.Published)

}
