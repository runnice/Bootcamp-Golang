package main

import "fmt"

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func checarSalario(salario int) (string, error) {
	if salario < 15000 {
		return "", &MyError{"erro: O salário digitado não está dentro do valor mínimo"}
	} else {
		return "Necessario pagamento de imposto", nil
	}

}

func checarError(salario int) {
	fmt.Printf("Salario: %d\n", salario)
	s, err := checarSalario(salario)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
func main() {

	salario := 10000
	checarError(salario)
	salario = 20000
	checarError(salario)

}
