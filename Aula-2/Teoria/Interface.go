package main

import (
	"fmt"
	"math"
)

//INTERFACE

// Uma interface é uma forma de definir métodos que devem ser usados, mas sem defini-los

// Interfaces são usadas para fornecer modularidade a um programa.

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Criamos uma função que imprime a área e o perímetro que geramos para esses objetos:

func detalhe(c circulo) {
	fmt.Println("Circulo de raio:", c)
	fmt.Println("Area:", c.area())
	fmt.Println("Perimetro:", c.perimetro())
}

// Se quiseremos gerar mais figuras geométricas? É aqui que a interface entra em jogo
//  A interface vai impedir que tenhamos que criar farias funções de detalhes.

type geometria interface {
	area() float64
	perimetro() float64
}

// Agora vamos criar uma estrutura chamada triangulo que terá os mesmos métodos:

type triangulo struct {
	largura, altura float64
}

func (t triangulo) area() float64 {
	return (t.largura * t.altura)
}

func (t triangulo) perimetro() float64 {
	return (2*t.largura + 2*t.altura)
}

// Vamos modificiar  nossa função de detalhes. para que ao invés de receber um círculo
// receba uma figura geométrica.
func detalheModificado(g geometria) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimetro:", g.perimetro())
}

// Desta forma podemos continuar adicionando figuras geométricas sem ter que modoficar
// nossa função de detalhes.

// No exemplo a seguir, criamos uma função que gera o objeto.

func novoCirculo(raio float64) circulo {
	return circulo{raio: raio}
}

// O que acontecerá se quisermos reutilizar nossa função para implementar várias figuras geométricas?

// Neste caso teremos que criar uma função que retorne uma interface.

const (
	tipoTriangulo = "triangulo"
	tipoCirculo   = "circulo"
)

// Vamos subistituir nossa função novoCirculo por novaFigura  e passamos 2 constantes que definimos para especificar qual é o objeto que geramos:

func novaFigura(geoTipo string, valores ...float64) geometria {
	switch geoTipo {
	case tipoTriangulo:
		return triangulo{largura: valores[0], altura: valores[1]}
	case tipoCirculo:
		return circulo{raio: valores[0]}
	}
	return nil
}

// INTERFACES VAZIAS //

// São as interfaces que não possuem métodos declarados

// Se são vazias... para que usamos?

// A utilizade dessas interfaces é nos fornecer um tipo de dado "coringa" ou seja,
// armazenar valores que são de um tipo de dado desconhecido,
// ou que podem variar de acordo com o fluxo do programa.

// COMO DECLARAMOS UMA VARIÁVEL DESSE TIPO?

// var minhaVariavel interface{}

// EXEMPLO DE USO DE INTERFACE VAZIA

type ListaHeterogenea struct {
	Data []interface{}
}

// linha 133 a 138

// TYPE ASSERTION - Declaração de tipo

// A Asserção de tipo fornece acesso ao tipo de dados exato que é abstraído por uma interface.

func main() {

	c := circulo{raio: 5}
	t := triangulo{largura: 10, altura: 5}

	detalheModificado(c)
	detalheModificado(t)

	c2 := novoCirculo(10)
	detalheModificado(c2)

	t3 := novaFigura(tipoTriangulo, 20, 5)
	detalheModificado(t3)

	c3 := novaFigura(tipoCirculo, 10)
	detalheModificado(c3)

	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "string")
	l.Data = append(l.Data, true)

	fmt.Println("%v\n", l.Data)

	var i interface{} = "Olá"

	str := i.(string) // checa se i é uma string = true
	fmt.Println(str)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

}
