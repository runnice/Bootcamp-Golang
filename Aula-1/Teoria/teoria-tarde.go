/* Operadores go
+  soma
-  subtração
* multiplicação
/ divisão
% módulo
++ incremento
-- decremento */

package main

import "fmt"

func main(){

	//%d = serve para concaternar strings no Println

	x, y := 10, 20

	//Operadores de 
	fmt.Println("x + y = %d\n", x + y)
	fmt.Println("x - y = %d\n", x - y)
	fmt.Println("x * y = %d\n", x * y)
	fmt.Println("x / y = %d\n", x / y)
	fmt.Println("x % y = %d\n", x % y)

	//Operadores de Atribuição
/* 	x=y 0 - Atribuição Simples
	x+=y - Operador de adição e atribição
	x-=y - Operador de subtração e atribuição
	x*=y - Operador de multiplicação e atribuição
	x/=y - Operador de divisão e atribuição
	x%=y - Operador de módulo e atribuição */

	// Operadores de Compração
	// x, y:= 10, 20
	// (x==y) é falso
	// (x!=y) é verdadeirox
	// (x>y) é falso
	// (x>=y) é falso
	// (x<y) é verdadeiro
	// (x<=y) é verdadeiro

	// Operadores lógicos
	// x, y, z:= 10, 20, 30
	// (x<y && x>) é falso
	// (x<y || x>z) é verdadeiro
	// !(x=y && x>z) é verdadeiro


	idade:= 10
	*p = idade
	fmt.Println(p)
	fmt.Println(&p)

	var a[2]string // declaração de Arrays
	var b[] string // declarar Array sem precisar declarar o tamanho do array.
	
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a, a[0], a[1])

	append(b, "oi")
	append(b, "tudobem?")

	myMap:= map[string]int{}
	fmt.Println(len(myMap))

	var students = map[string]int{"João":22, "Pedro": 30} // Sintaxe quer um mapa de String e o segundo valor um Inteiro.
	fmt.Println(students["João"])

	var estudantes := make(map[string]int)
	fmt.Println(estudantes)
	fmt.Println(len(estudantes))

	students := make(map[string]interface{}) // Quando não sabe o segundo parâmetro.
	students := make (map[string]any) // Quando não sabe o que vai receber como parâmetro.

	delete(students, "João") //Delete palavra reservada para deletar atributos, no caso deletará João:22 do map.

	Tipos de For:

	Standard
	
	for i:=0; i<100;i++
	sum+=sum
	
	For Range

	frutas:=[]string{"banana", "maçã", "laranja"}
	for key, value := range frutas{
		fmt.Println(key, value)
	}

	Se não quiser o primeiro parâmetro, podemos utilizar o _ que o Go entenderá.
	
	frutas:=[]string{"banana", "maçã", "laranja"}
	for _, value := range frutas{
		fmt.Println(value)
	}


	Loop Infinito 

	Vai rodar para sempre

	sum:=0
	for{
		fmt.Println(sum++)

	}

	Loop While
	
	for sum<10{
		sum += sum
	}

	Quebra de laço

	for{
		sum++
		if sum >= 1000{
			break
		}
	}
	fmt.Println(sum)


	Pular para próxima interação utiliza-se o continue
	
	for i:=0; i<10;i++{
		if i%2 ==0{
			continue
		}
		fmt.println(i, "is odd")
	}

	Declaração Curta
	if var declaração; condição{
		bloco de código
	}
	if result = sum(5,5); result < 10{
		fmt.Println("Resultado é menor que 10")
	}

	if result1, result2 := sum(5,5); result2<10 && result1 ==5{
		fmt.Println("Resultado é menor que 10")
	}else {
		fmt.Println("Resultado é maior que 10")
	}

	Switch Básico
	dias:=1
	switch dias{
	case 0:{
		fmt.Println("Segunda")
	case 1:{
		fmt.Println("Terça")
	case 2:{
		fmt
	}
	}}




	
	
}
}