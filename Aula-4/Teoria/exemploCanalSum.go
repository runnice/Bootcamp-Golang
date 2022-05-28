package main

import (
	"fmt"
)

// <-chan tudo que vou usar na função indepente do  resultado
// out chan<- tudo que vou jogar pra fora

func main() {

	grades := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	sumChan := make(chan int)
	avgChan := make(chan int)

	go sum(grades, sumChan)
	go avg(sumChan, avgChan, len(grades))

	printAvg(avgChan)

}

func sum(grades []int, out chan<- int) {
	var total int
	for _, grade := range grades {
		total += grade
	}

	out <- total
	close(out)
}

func avg(in <-chan int, out chan<- int, lenght int) {
	var avgGrades int
	for v := range in {
		avgGrades = v / lenght
	}
	out <- avgGrades
	close(out)
}

func printAvg(out <-chan int) {
	for v := range out {
		fmt.Println(v)
	}

}
