package main

import (
	"errors"
	"fmt"
)

type Funcionario struct {
	horasTrabalhadas int
	valorDaHora      float64
}

func salarioMensalPorHoras(horasTrabalhada int, valorDaHora float64) (float64, error) {
	if horasTrabalhada < 80 {
		return 0, errors.New("Erro: O trabalhador não pode ter trabalhado menos que 80 horas no mês")
	}
	salario := float64(horasTrabalhada) * valorDaHora
	if salario >= 20000 {
		salario *= 0.9
	}
	return salario, nil

}
func main() {

	gustaFuryBoy := Funcionario{horasTrabalhadas: 50, valorDaHora: 200}
	funcionario1 := Funcionario{horasTrabalhadas: 80, valorDaHora: 100}
	funcionario2 := Funcionario{horasTrabalhadas: 70, valorDaHora: 100}

	salarioGustaFuryBoy, err := salarioMensalPorHoras(gustaFuryBoy.horasTrabalhadas, gustaFuryBoy.valorDaHora)
	if err != nil {
		e1 := fmt.Errorf("Erro: %w", err)
		e2 := fmt.Errorf("e1: %w", e1)
		fmt.Println(errors.Unwrap(e2))
	} else {
		fmt.Printf("Salario do funcionario GustaFuryBoy: %.2f\n", salarioGustaFuryBoy)
	}

	funcionarioHumilde, err := salarioMensalPorHoras(funcionario1.horasTrabalhadas, funcionario1.valorDaHora)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Salario do funcionario funcioanrioHumilde: %.2f\n", funcionarioHumilde)
	}

	funcionarioComFilhoEmCasa, err := salarioMensalPorHoras(funcionario2.horasTrabalhadas, funcionario2.valorDaHora)
	if err != nil {
		e1 := fmt.Errorf("Erro: %w", err)
		e2 := fmt.Errorf("e1: %w", e1)
		fmt.Println(errors.Unwrap(e2))
	} else {
		fmt.Printf("Salario do funcionario funcionarioComFilhoEmCasa: %.2f\n", funcionarioComFilhoEmCasa)
	}

}
