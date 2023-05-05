package main

import (
	"fmt"
)

func g1(dataChan chan<- int, doneChan chan<- bool, data int) {
	dataChan <- data
	doneChan <- true
}

func g2(dataChan chan<- int, doneChan chan<- bool, data int) {
	dataChan <- data
	doneChan <- true
}

func g3(dataChan1 <-chan int, dataChan2 <-chan int, sumChan chan<- int) {
	sum := 0
	for i := 0; i < 2; i++ {
		select {
		case data := <-dataChan1:
			sum += data
		case data := <-dataChan2:
			sum += data
		}
	}
	sumChan <- sum
}

func main() {
	dataChan1 := make(chan int)
	doneChan1 := make(chan bool)
	dataChan2 := make(chan int)
	doneChan2 := make(chan bool)
	sumChan := make(chan int)

	go g1(dataChan1, doneChan1, 1)
	go g2(dataChan2, doneChan2, 8)

	go func() {
		<-doneChan1
		<-doneChan2
		close(dataChan1)
		close(dataChan2)
	}()

	go g3(dataChan1, dataChan2, sumChan)

	sum := <-sumChan
	fmt.Println("Sum:", sum)
}
