package main

import (
	"fmt"
)

func checarSalario3(salario int) (string, error) {
	if salario < 15000 {
		return "", fmt.Errorf("erro: O mínimo tributável é 15.000 e o salário informado é: %d", salario)
	} else {
		return "Necessario pagamento de imposto", nil
	}

}

func checarError3(salario int) {
	fmt.Printf("Salario: %d\n", salario)
	s, err := checarSalario3(salario)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
func main() {

	salario3 := 10000
	checarError3(salario3)
	salario3 = 20000
	checarError3(salario3)

}
