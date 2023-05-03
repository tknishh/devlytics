package main

import "fmt"

func main() {
	dataChan1 := make(chan int)
	dataChan2 := make(chan int)
	resultChan := make(chan int)

	go g1(dataChan1)
	go g2(dataChan2)

	go g3(dataChan1, dataChan2, resultChan)

	// Send integer to g1 and g2
	dataChan1 <- 10
	dataChan2 <- 20

	// Wait for result from g3
	result := <-resultChan

	fmt.Println("Result:", result)
}

func g1(dataChan chan<- int) {
	for {
		data := <-dataChan
		fmt.Println("g1 received data:", data)
		dataChan <- data
	}
}

func g2(dataChan chan<- int) {
	for {
		data := <-dataChan
		fmt.Println("g2 received data:", data)
		dataChan <- data
	}
}

func g3(dataChan1 <-chan int, dataChan2 <-chan int, resultChan chan<- int) {
	for {
		data1 := <-dataChan1
		data2 := <-dataChan2
		result := data1 + data2
		fmt.Printf("g3 received %d and %d, adding...\n", data1, data2)
		resultChan <- result
	}
}
