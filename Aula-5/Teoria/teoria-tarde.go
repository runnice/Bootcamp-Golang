package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ERROS NA LINGUAGEM GO

type MyError struct {
	status int
	msg    string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

// Função recebe um erro e retorna dois valores

func myCustomErrorTeste(status int) (int, error) {
	if status >= http.StatusBadRequest {
		return http.StatusInternalServerError, &MyError{
			status: status,
			msg:    "Algo deu error",
		}
	}
	return http.StatusOK, nil
}

// errors.New()

// errors.New()

func sayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("o nome está vazio")
	}
	return "Hello, " + name, nil
}

// fmt.Errorf() - Pacote para formatar o Error
// fmt.Errorf(" mensagem + marcador %d", d, 400)

// Por convenção mensagem de error é tudo minúsculo e sem pontuação.

// Outras funções do pacote error New(), Is(), As(), Unwrap()

// New - passa o novo error
// As - passa o tipo do error se encontrar correspondência retorna true
// Is - passa 2 errors e compara se eles são iguais.

// var err1 = errors.New("error número 1")

// func x() error{
// 	return fmt.Errorf("informação extra do erro: %w", err1)
// }
// e:=x()
// coincidence :=error.Is(e, err2)
// fmt.Println(coincidence)

// Unwrap()
type errorTwo struct{}

func (e errorTwo) Error() string {
	return "error two happened"
}

//Papel do multi-retorno no tratamentos de erros?
// Go perimite que as funções retornem mais de um valor. Entre os valores retornados

//Sintaxis

func main() {
	status, err := myCustomErrorTeste(300)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d, funciona!", status)

	name, err := sayHello("Cyro")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(name)

	e := &MyError{1, "404"}
	var err2 *MyError
	isMyError := errors.As(e, &err2)

	fmt.Println(isMyError)

	// Unwrap

	e2 := errorTwo{}
	e1 := fmt.Errorf("e2: %w", e2)
	fmt.Println(errors.Unwrap(e1))
	fmt.Println(errors.Unwrap(e2))
}
