package main

import (
	"errors"
	"fmt"
)

func checarSalario2(salario int) (string, error) {
	if salario < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	} else {
		return "Necessario pagamento de imposto", nil
	}

}

func checarError2(salario int) {
	fmt.Printf("Salario: %d\n", salario)
	s, err := checarSalario2(salario)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
func main() {

	salario2 := 10000
	checarError2(salario2)
	salario2 = 20000
	checarError2(salario2)

}
