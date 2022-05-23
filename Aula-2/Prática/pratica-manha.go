package main

import (
	"errors"
	"fmt"
)

func calcImposto(salario, imposto float64) float64 {
	imposto = imposto / 100
	return salario * imposto

}

func calculaMedia(notas ...int) (float64, error) {
	res, qntdNotas := 0, 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Não é possível calcular a média com valores negativos")
		}
		qntdNotas++
		res += nota

	}
	return float64(res / qntdNotas), nil
}

func calculoSalario(categoria string, horas float64) float64 {
	impostoCategoriaA := 0.5
	impostoCategoriaB := 0.2
	salarioCategoriaC := 1000.
	salarioCategoriaB := 1500.
	salarioCategoriaA := 3000.

	switch categoria {
	case "C":
		salario := &salarioCategoriaC
		return *salario * horas
	case "B":
		salario := &salarioCategoriaB
		if horas > 160 {
			return ((*salario + (*salario * impostoCategoriaB)) * horas)
		} else {
			return *salario * horas
		}
	case "A":
		salario := salarioCategoriaA
		if horas > 160 {
			return (salario + (salario * impostoCategoriaA)) * horas
		} else {
			return salario * horas
		}
	default:
		fmt.Printf("Categoria inválida")
	}
	return 0
}

func minimum(valores ...float64) (float64, error) {
	minimo := valores[0]
	if len(valores) == 0 {
		return 0, errors.New("Não é possível calcular o mínimo com valores vazios")
	}
	for i := 0; i < len(valores); i++ {
		if valores[i] < minimo {
			minimo = valores[i]
		}
	}
	return minimo, nil
}

func maximum(valores ...float64) (float64, error) {
	maximo := valores[0]
	if len(valores) == 0 {
		return 0, errors.New("Não é possível calcular o máximo com valores vazios")
	}

	for i := 0; i < len(valores); i++ {
		if valores[i] > maximo {
			maximo = valores[i]
		}
	}
	return maximo, nil
}

func average(valores ...float64) (float64, error) {
	avg := 0.
	count := 0.

	if len(valores) == 0 {
		return 0, errors.New("Não é possível calcular a média com valores vazios")
	}
	for _, valor := range valores {
		avg += valor
		count++
	}
	return avg / count, nil
}

func operation(operacao string) (func(valores ...float64) (float64, error), error) {
	switch operacao {
	case "minimum":
		return minimum, nil
	case "average":
		return average, nil
	case "maximum":
		return maximum, nil
	default:
		return nil, errors.New("Operação inválida")
	}
}

func doguinho(qntd float64) float64 {
	return qntd * 10000
}
func cat(qntd float64) float64 {
	return qntd * 5000
}
func hamster(qntd float64) float64 {
	return qntd * 250
}
func tarantula(qntd float64) float64 {
	return qntd * 150
}

func Animal(animal string) (func(qntd float64) float64, error) {
	switch animal {
	case "dog":
		return doguinho, nil
	case "cat":
		return cat, nil
	case "hamster":
		return hamster, nil
	case "tarantula":
		return tarantula, nil
	default:
		return nil, errors.New("Animal inválido")
	}
}

func main() {

	const (
		salario1 = 50000
		salario2 = 100000
		imposto1 = 17
		imposto2 = 10
	)

	//Exercício 1:

	var employees = map[float64]float64{salario1: imposto1, salario2: imposto2}
	for key, value := range employees {
		fmt.Println("Salario do funcionário é : R$", key, "Imposto:", value, "%")
		fmt.Println("Imposto do funcionário é: R$", calcImposto(key, value))
	}

	//Exercício 2:

	media, err := calculaMedia(10, 20, 30, 40, 50)
	if err != nil {
		fmt.Println("Sua operação falhou", err)
	} else {
		fmt.Println("A média do aluno é:", media)
	}

	media2, err := calculaMedia(-2, -4, -6, -8, 10)
	if err != nil {
		fmt.Println("Sua operação falhou", err)
	} else {
		fmt.Println("A média do aluno é:", media2)
	}

	//Exercício 3:

	fmt.Println(calculoSalario("C", 162))
	fmt.Println(calculoSalario("B", 176))
	fmt.Println(calculoSalario("A", 172))
	fmt.Println(calculoSalario("D", 172))

	// Exercício 4:
	const (
		minimum   = "minimum"
		average   = "average"
		maximum   = "maximum"
		dog       = "dog"
		cat       = "cat"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	minhaFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	minValue, err := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue, err := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, err := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)

	//Exercício 5

	var amount float64
	funcDoguinho, err := Animal(dog)
	funcGatinho, err := Animal(cat)
	funcHamster, err := Animal(hamster)
	funcTarantula, err := Animal(tarantula)

	amount += funcDoguinho(1)
	amount += funcGatinho(1)
	amount += funcHamster(1)
	amount += funcTarantula(1)

	kg := amount / 1000

	fmt.Println("Total de alimentos:", kg, "kg")

}
