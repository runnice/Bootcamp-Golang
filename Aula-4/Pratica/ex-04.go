package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func insertionSort(items []int, out chan<- time.Duration) {
	tempo := time.Now()

	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	tempoDepois := time.Now()
	out <- tempoDepois.Sub(tempo)
}

func selectionSort(items []int, out chan<- time.Duration) {
	tempo := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	tempoDepois := time.Now()
	out <- tempoDepois.Sub(tempo)
}

func groupSort(items []int, out chan<- time.Duration) {
	tempo := time.Now()
	sort.Ints(items)
	tempoDepois := time.Now()
	out <- tempoDepois.Sub(tempo)
}

func main() {
	cemNumeros := rand.Perm(100)
	milNumeros := rand.Perm(1000)
	dezMilNumeros := rand.Perm(10000)

	chInsertSort := make(chan time.Duration)
	chSelectSort := make(chan time.Duration)
	chGroupSort := make(chan time.Duration)

	fmt.Println("Executando Insertion Sort com 100 elementos")
	go insertionSort(cemNumeros, chInsertSort)
	t1 := <-chInsertSort
	fmt.Println(" Tempo Insertion Sort: ", t1)
	fmt.Println(" Organizando os números: ")
	go groupSort(cemNumeros, chGroupSort)
	ts := <-chGroupSort
	fmt.Println("tempo", ts)

	fmt.Println("Executando Selection Sort com 100 elementos")
	go selectionSort(cemNumeros, chSelectSort)
	t2 := <-chSelectSort
	fmt.Println("Tempo Selection Sorte: ", t2)

	fmt.Println("Executando Group Sort com 100 elementos")
	go insertionSort(milNumeros, chInsertSort)
	fmt.Println("Insertion Sort: ", t2)
	go insertionSort(dezMilNumeros, chInsertSort)

	go selectionSort(milNumeros, chSelectSort)
	go selectionSort(dezMilNumeros, chSelectSort)

	go groupSort(cemNumeros, chGroupSort)
	go groupSort(milNumeros, chGroupSort)
	go groupSort(dezMilNumeros, chGroupSort)
}
