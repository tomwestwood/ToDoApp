package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
)

func main() {
	raceCorrect_InOrder(os.Stdout)
}


func raceCorrect_InOrder(writer io.Writer) {
	var wg sync.WaitGroup
    
	oddStream := newIncrementerStream(&wg, 1)	
	evenStream := newIncrementerStream(&wg, 2)

	wg.Wait()
	var numbers []int

	numbers = appendStreamToList(oddStream, numbers)
	numbers = appendStreamToList(evenStream, numbers)

	sort.Ints(numbers)
	for _, result := range numbers {		
		fmt.Fprintln(writer, result)
	}
	
}

func appendStreamToList(stream chan int, numbers []int) []int {
	for result := range stream {
		numbers = append(numbers, result)
	}
	return numbers
}

func newIncrementerStream(wg *sync.WaitGroup, startOn int) chan int {
	incrementerStream := make(chan int)
	go func() {
		wg.Add(1)
		defer wg.Done()
		defer close(incrementerStream)
		for i := startOn; i <= 10; i += 2 {
			incrementerStream <- i
		}
	}()
	return incrementerStream
}