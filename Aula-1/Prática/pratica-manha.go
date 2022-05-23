package main

import "fmt"

func main() {

	//Exercício 1
	nome, idade := "Vinicius", 35
	fmt.Println(nome, idade)

	//Exercício 2
	var temperature float32 = 29.5
	var moisture float32 = 0.5
	var atmosphericPressure float32 = 1015

	fmt.Println("A temperatura atual é", temperature)
	fmt.Println("A umidade atual é", moisture)
	fmt.Println("A pressão atmosférica atual é", atmosphericPressure, "hPa")

	//Exercício 3

	//var 1nome string                // Variável começando com número. Correção: var nome1 string
	//var sobrenome string            // correto
	//var int idade	                  // A tipagem tem que vir depois do nome da variável. Correção: é var idade int ou até mesmo o uint já que idade não fica negativo.
	//1sobrenome1 := 6                // Variável começando com número. Correção: sobrenome1 := 6
	//var licenca_para_dirigir = true // variável separada com underline, o ideal é a utilização do camelCase. var licensaParaDirigir
	//var estatura da pessoa int      // variável com espaço. deve ser corrigida para camelCase. var estaturaDaPessoa int
	//quantidadeDeFilhos := 2         // correto

	//Exericio 4

	var sobrenome string = "Silva"
	var idade uint8 = 25
	boolean := false
	var salario float32 = 4585.50
	var nome string = "Fellipe"

}
