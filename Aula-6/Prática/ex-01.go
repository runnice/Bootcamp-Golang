package main

import (
	"fmt"
	"os"
)

func main() {

	arq, err := os.Open("customers.txt")
	if err != nil {
		fmt.Println("o arquivo indicado não foi encontrado ou está danificado")
		panic(err)
	}
	defer arq.Close()

}
