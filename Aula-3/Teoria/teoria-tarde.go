package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Compreender pacotes fmt, os e io
// pkg.go.dev/fmt

//PACKAGE FMT

// Nos permite trabalhar com formatação

// \n quebra de linha
// \\ barra invertida
// \t tabulação horizontal
// \v tabuação vertical

// Verbos de impressões mais utilizados

// %v valor no formato padrão
// %T tipo de dado do valor a imprimir
// %t bool
// %s string
// %f ponto flutuante
// %d decimal
// %b binário
// %o octal
// %x hexadecimal
// %c caractere

// .Print() toma como paramêtro qualquer tipo, porém não consegue utilizar os verbos.
// .Println() já coloca um espaço no fim da linha.
// . Printf() Consegue formatar usando os verbos.

// .Sprint() você consegue armazenar numa variável e imprimir depois
// .Sprintf() você imprimir o Sprint formatado

// PACKAGE OS

// O pacote "os" serve para executar funcionalidades do sistema operacional
// Conhecer infos de um arquivo, como exemplo:
// Nome, Tamanho, Data
// Variaveis de ambiente:
// ler, escrever, apagar, renomear, criar, mover, copiar, etc.

// .Setenv(Setar variável de ambiente)
// .Getenv(Capturar a envi setada)

// .LookupEnv - Equivalente a Getenv porém retorn 2 valores. O nome da variável  e um booleano se existe ou não.

// .ReadDir() -- Lê o diretório nomeado, retornando todas as suas entradas de diretório
// ordendadas pelo nome

// .ReadFile() - recebe como parâmetro o endereço e o nome do arquivo e retorna o conteudo em bytes do arquivo.

// .WriteFile() - Recebe

// .Copy()  Pega um leitor e copia dados de origem para o destino.

// .WriteString() - Escreve uma string no terminal.

func main() {

	nome, idade := "Vinicius", 30

	// Função Print reebe um erro tmb. Exemplo de como suar abaixo:
	ok, err := fmt.Printf("Meu nome é : %s. Tenho %d anos.", nome, idade)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	fmt.Println(ok)

	fmt.Printf("%10.2f\n", 3.141592653589793) // 10 espaços totais e 2 depois do ponto
	fmt.Printf("%.2f\n", 3.141592653589793)   // Qualquer qntd a esquerda e 2 depois do ponto

	res := fmt.Sprint("Meu nome é: ", nome, " e tenho ", idade, " anos.\n")
	fmt.Print(res)

	res = fmt.Sprintf("Meu nome é: %s e tenho %d anos.", nome, idade)
	fmt.Println(res)

	//setar variável de ambiente

	err2 := os.Setenv("Name", "Gopher")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(err2)

	//capturando a env setada
	value := os.Getenv("Name")
	if err != nil {
		log.Println("ocorreu um erro ao setar a variável de ambiente", err)

	}
	log.Println(value)

	// log.SetFlags(0) - Se não quiser horário no log

	// uso do LookupEnv
	value2, existe := os.LookupEnv("Name")
	if existe {
		log.Println("a variável existe: ", value)
	} else {
		log.Println("a variável não existe")
	}
	log.Println(value2)

	// Uso so .ReadDir()
	files, err := os.ReadDir(".")
	if err != nil {
		log.Println("Ocorreu um erro ao ler o diretório", err)
	}

	// Para imprimir os arquivosa do files, é necessário um for.
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Uso do .ReadFile() - ele retornará um array de bytes e precisaremos transformar em string.

	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Println("Erro ao ler o arquivo", err)
	}

	log.Println(string(data))
	log.Printf("%s", data) // Faz o mesmo que a string(data) imprime

	d1 := []byte("\nOla mundo, muundão!")
	err3 := os.WriteFile("test.txt", d1, 0664)
	if err3 != nil {
		log.Println("erro ao escrever o arquivo", err)
	}

	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(string(d1)); err != nil {
		panic(err)
	}

	// Uso da função Copy

	r := strings.NewReader(" Some io.reader stream to be read\n")
	_, err4 := io.Copy(os.Stdout, r)
	if err4 != nil {
		log.Println(err)
	}
	log.Println(r)

	// Uso de WriteString

	okay, err5 := io.WriteString(os.Stdout, "Hello, world!\n")
	if err5 != nil {
		log.Println(err)
	}
	log.Println(string(okay))
}
