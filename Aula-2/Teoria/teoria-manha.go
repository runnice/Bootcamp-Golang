package main

import (
	"errors"
	"fmt"
	"log"
)

//Função Void função sem retorno
func logger() {
	log.Println("loggando...") // Println é um método do pacote log
	fmt.Println("logando....") // Println é um método do pacote fmt
}

//Parametros de uma função

//Função com retorno
func soma(x, y float32) float32 { //o que vem depois do parentese é o retorno
	return x + y
}

//Toda vez que a função tiver retorno, ela pode ser armazenada em uma variável

// minhaSoma := soma(10, 15)
// fmt.Println(minhaSoma)

// Exemplo mais complexo
const (
	Soma          = "+"
	Subtracao     = "-"
	Multiplicacao = "*"
	Divisao       = "/"
)

func operacaoAritimetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Soma:
		return valor1 + valor2
	case Subtracao:
		return valor1 - valor2
	case Multiplicacao:
		return valor1 * valor2
	case Divisao:
		return valor1 / valor2
	default:
		return 0

	}
}

// Exemplo de Enum, usamos o iota
// const (
//	PRETO = iota // este é o 0
//	BRANCO // este é o 1
// AZUL // este é o 2
//)

// Escopo = Até onde as chaves conseguem enxergar

// Ellipsis (...)
// Ela também é conhecida como função variádica ou função de múltiplos retornos
// coloca as reticiadas (...) no final da função para indicar que ela pode retornar vários valores

//Por termos usado o valores ...float64, o valores vira um array de float64
//Para retornar mais de um valor, usamos o (...) Ellipsis

func soma2(valores ...float64) float64 {
	var resultado float64
	for _, valor := range valores {
		resultado += valor
	}
	return resultado
}

// Funções para realizar cada operação

func opSoma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}
func opSubtracao(valor1, valor2 float64) float64 {
	return valor1 - valor2
}
func opMultiplicacao(valor1, valor2 float64) float64 {
	return valor1 * valor2
}
func opDivisao(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

// Atibuindo uma função dentro de outra função
// Função encarregada de receber a operação a ser realizada e os valores aos quais a operação será aplicada.
// Para cada operaão, criamos uma função que recebe os valores e a função que executa aquele operador.

func operacaoAritimetica2(operador string, valores ...float64) float64 {
	switch operador {
	case Soma:
		return retornoOperacao(valores, opSoma)
	case Subtracao:
		return retornoOperacao(valores, opSubtracao)
	case Multiplicacao:
		return retornoOperacao(valores, opMultiplicacao)
	case Divisao:
		return retornoOperacao(valores, opDivisao)

	}
	return 0
}

// Criamos uma função que se encarrega de retornar as operações.
// []float64 - criarar um array de valores.
func retornoOperacao(valores []float64, operacao func(valor1 float64, valor2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacao(resultado, valor)
		}
	}
	return resultado
}

// Função com Multiplos retornos.

// Temos que indicar os tipos de dados dos valores que serão retornados, separados por vírgulas e entre parênteses.

func operations(valor1, valor2 float64) (float64, float64, float64, float64) {
	soma := valor1 + valor2
	subtra := valor1 - valor2
	mult := valor1 * valor2

	var div float64

	if valor2 != 0 {
		div = valor1 / valor2
	}

	return soma, subtra, mult, div
}

//Ao chamar nossa função, devemos receber todos os valores que ela retorna.

//Package Errors

// Em GO o retorno multivalor é normalmente utilizado quando precisamos retornar um valor e um erro, e precisamos validar se ocorreu um erro ou não
// Para isso vamos fazer um exemplo de divisão e ela retorna um erro caso o divisor seja zero.

func divisao(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("O divisor não pode ser zero")
	}
	return dividendo / divisor, nil
}

// Retorno de valores Nomeados:
// Também podemos retornar valores nomeados. Para isso devemos definir na função não apenas o tipo de dados a ser retornado mas também o nome da variável

func operaciones(x, y float64) (soma float64, sub float64, mulit float64, div float64) { // Podemos nomear as variáveis juntas e só tipar no final, se todas compartilharem do mesmo tipo.
	soma = x + y
	sub = x - y
	mulit = x * y
	div = x / y
	return soma, sub, mulit, div
}

// Retorno de Função: Como uma função retorna outra função.
func operacaoAritimeticaRecebendoFuncao(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Soma":
		return opSoma
	case "Subtracao":
		return opSubtracao
	case "Multiplicacao":
		return opMultiplicacao
	case "Divisao":
		return opDivisao
	}
	return nil
}

// Va para a linha 219 para ver o retorno de função.

func main() {

	logger()
	minhaSoma := soma(10, 15)
	fmt.Println(soma(5, 5), minhaSoma)
	fmt.Println(soma2(5, 5))

	fmt.Println("Operacao Aritimética Soma: ", operacaoAritimetica(10, 5, Soma))
	fmt.Println("Operacao Aritimética Subtração: ", operacaoAritimetica(10, 5, Subtracao))
	fmt.Println("Operacao Aritimética Multiplicação: ", operacaoAritimetica(10, 5, Multiplicacao))
	fmt.Println("Operacao Aritimética Divisão: ", operacaoAritimetica(10, 5, Divisao))

	fmt.Println(operacaoAritimetica2(Soma, 2, 3, 4, 5, 6))

	//Ao chamar nossa função operations, devemos receber todos os valores que ela retorna.

	s, r, m, d := operations(10, 5)

	fmt.Println("Soma: \t\t", s)
	fmt.Println("Subtração: \t\t", r)
	fmt.Println("Multiplicação: \t\t", m)
	fmt.Println("Divisão: \t\t", d)

	// ERRORS - Package Errors
	// Executando a função divisao e verificando se ocorreu um erro.
	// Exemplo de como tratar o erro:

	res, err := divisao(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	somas, subtras, multis, diviss := operaciones(10, 20)
	fmt.Println("Operações", somas, subtras, multis, diviss)

	//Retorno de função que recebe função:
	oper := operacaoAritimeticaRecebendoFuncao("Soma")
	resultado := oper(2, 5)
	fmt.Println(resultado)

}
