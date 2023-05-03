package main

import (
	"fmt"
	"sync"
)

func g1(dataChan chan<- int, data int) {
	dataChan <- data
}

func g2(dataChan chan<- int, data int) {
    dataChan <- data
}

func g3(dataChan <-chan int, sumChan chan<- int) {
    sum := 0
    for data := range dataChan {
        sum += data
    }
    sumChan <- sum
}

func main() {
    dataChan1 := make(chan int)
    dataChan2 := make(chan int)
    sumChan := make(chan int)

    go g1(dataChan1, 1)
    go g2(dataChan2, 2)
    go g3(merge(dataChan1, dataChan2), sumChan)

    sum := <-sumChan
    fmt.Println("Sum:", sum)
}

// merge merges two channels into one
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start a goroutine for each input channel to read values
    // and send them to the output channel
    output := func(ch <-chan int) {
        defer wg.Done()
        for n := range ch {
            out <- n
        }
    }
    wg.Add(len(channels))
    for _, ch := range channels {
        go output(ch)
    }

    // Wait for all input goroutines to complete
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}