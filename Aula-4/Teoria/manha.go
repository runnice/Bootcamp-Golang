package main

// PONTEIRO

// um ponteiro é um endereço de memória que referencia outro valor

// * para criar variável do tipo pointeiro
// * para visualizar o valor da posição da memória
// & para atribuir um valor ao ponteiro
// Exemplo de função Ponteiro:

//Aumentar recebe um ponteiro do tipo inteiro
import (
	"fmt"
	"time"
)

func Aumentar(p *int) {
	*p++
}

//Simultaneidade e Paralelismo

// -- Simultaneidade ou Concorrência
// A simultanedade é o ator de iniciar, executar e concluir duas ou mais tarefas em preíodos de tempos escalonados.
// Não significa necessariamente que ambos estão sendo executados ao mesmo tempo

// -- Paralelismo
// Paralelismo implica que duas ou mais tarefas sejam executados ao mesmo tempo

// Go Routines
// Eles são solução oferecidas pela GO para implementar a simultaneidade
// NÃO SÃO THREADS! Go routines são gerenciadas pelo GO Runtime e seu agendador, não pelo SO

// Quando executamos uma função Go routines, sua execução não bloqueará a continuação da função que a chamou

func processar(i int) {
	fmt.Println(i, "- Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")

}

// Um canal nos permitirá enviar valores para as Go Routines
// para defindir um canal do tipo inteiro fazemos da seguinte forma:

// c:= make(chan(int))

func processarComCanal(i int, c chan int) {
	fmt.Println(i, "- Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i // c - seta e o valor que você quer enviar.
}

func main() {
	// Criando ponteiros com new
	var p1 *int
	var p2 = new(int)

	fmt.Println(p1)
	fmt.Println(p2)

	// OPERADOR DE DESREFERIANÇÃO
	var v int = 19
	var p *int = &v

	fmt.Printf("O ponteiro p refere-se ao endereço de memória %v \n", p)
	fmt.Printf("Ao desreferenciar o ponteiro p, obtenho o valor: %d \n", *p)

	// Função Aumentar com ponteiro
	Aumentar(&v)
	fmt.Println("O valor de v é:", v)

	//função processar, simulando um processamento
	// processar(3)

	for i := 0; i < 10; i++ {
		processar(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Programa encerrado")

	// Como executar uma função Go routine (Simultanea)?
	// Para executar uma função Go routine, utilizamos o comando go

	for i := 0; i < 10; i++ {
		go processar(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Programa encerrado")

	// Criando um canal
	canal := make(chan int)
	// Função Processar
	processarComCanal(1, canal)
	variavel := <-canal
	fmt.Println("O valor recebido do canal é:", variavel)
	//close(c) - Fecha o canal, se tentarmos usar abaixo ele não será possível.

	//Travar operação no banco para uma única operação
	// var wg sync.WaitGroup
	// wg.Wait()
	// defer wg.Done()

}
