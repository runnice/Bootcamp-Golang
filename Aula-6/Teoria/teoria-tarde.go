package main

import (
	"context"
	"fmt"
)

//PANIC
// Encerrar o programa com indicação que algo deu erro inesperadamente
// panic("Algo deu errado")

// Próprio GO utiliza PANIC por debaixo dos panos.
// Trabalhar com canais que já foram fechados ou consumidos
// Exceder a capacidade de indexação de um array
// Tentar trabalhar com

// Index Out of Bound Panics

// Tenta acessar uma posição que não existe no Array

// Receptores nulos

// Ocorre quando tentamos acessar um pontiero que não está apontando pra nenhum endereço de memória

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, ": Au Au")
}

// DEFER AND RECOVERY

// Conseguimos tratar um panic usando  defer e recover podemos controlar os efeitos de um panic e impedir que nosso programa termine
// Detalhe: Defer e revoer são funções nativas da linguagem especificamente projetadas para previnir ou controlar a natureza de um panic.

// DEFER
// É utilizado para deixar a execução tardia
// DEFER FUNCIONA COMO PILHA! Se tiver vários o primeiro a entrar é o último a sair

// log.Fatal() também funciona para matar o programa.

// Função Recover

// Conseguimos retornar mesmo tendo um panic.
// O programa continuará rodando.

func isEven(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if (num % 2) != 0 {
		panic("nao é par")
	}
	fmt.Println(num, " é par")
}

// Função que poderiamos chamar no isEven ao invés da função anônima.

func debugger() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}

// PACKAGE CONTEXT

// O pacote de context é uma das libs mais uteis que o o GO oferece
// Serve pra definir um vonexto que pode

// CONTEXT BACNGROUND

func hello() {
	fmt.Println("hello")
}
func main() {

	animals := []string{
		"vaca",
		"cachorro",
		"gato",
	}
	// Teria que passar o -1 no final do array na função len para não retornar PANIC
	fmt.Println("O animal que voa é o: ", animals[len(animals)])

	s := &Dog{"Sammy"} // S recebeu uma estrturua e uma estrutura sempre é um ponteiro.
	s = nil            //Aqui dizemos que S agora é nulo e um ponteiro.
	s.WoofWoof()       // Tentamos imprimir S que não aponta pra lugar nenhum

	// Usando Defer, o Hello será impresso depois do World.
	defer hello()
	fmt.Println("world")

	// Recover
	isEven(2)
	type contextKey string
	const ContextKeyID contextKey = iota
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKeyID, "Hello")
	fmt.Println(ctx)
}
