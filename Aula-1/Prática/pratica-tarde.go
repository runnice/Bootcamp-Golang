package main

import "fmt"

func main() {
	//Exercício 1

	var name string
	var idade, mes, count uint8
	var salario, tempoAtividade float64

	fmt.Println("Bem vindo ao Bytebank!")
	fmt.Println("Digite seu primeiro nome: ")
	fmt.Scan(&name)

	fmt.Println("Seu nome tem", len(name), "letras")

	fmt.Println("Vou soletrar seu nome")

	for _, letra := range name {
		fmt.Println(string(letra))
	}

	//Exercício 2

	fmt.Println("Por favor, me informe alguns dados.")

	fmt.Println("\nQuantos anos você tem?")
	fmt.Scan(&idade)
	fmt.Println("\nA quantos anos você trabalha")
	fmt.Scan(&tempoAtividade)
	fmt.Println("\nQual o seu rendimento?")
	fmt.Scan(&salario)

	switch {
	case idade < 22 || tempoAtividade < 1:
		fmt.Println(name, ", seu emprestimo não pode ser realizado")
	case idade > 22 && salario >= 100000:
		fmt.Println(name, ", seu emprestimo pode ser realizado com taxa 0!")
	case idade >= 22 && tempoAtividade >= 1:
		fmt.Println(name, ", seu emprestimo pode ser realizado")
	}

	//Exercicio 3

	meses := map[uint8]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	fmt.Println("\nQual o mês que você quer saber?")
	fmt.Scan(&mes)
	fmt.Println("\nO mês ", mes, " corresponde ao mês de ", meses[mes])

	//Exercicio 4

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("\nO funcionário Bejamin tem ", employees["Benjamin"], " anos")
	fmt.Println("")

	for key, value := range employees {
		if value < 21 {
			continue

		}
		fmt.Println("O funcionário ", key, "tem ", value, "anos")
		count++
	}
	fmt.Printf("\nForam encontrados %d funcionários com mais de 21 anos", count)

	delete(employees, "Pedro")
	fmt.Println("\nO funcionário Pedro foi removido do map")

}
