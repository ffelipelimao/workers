package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	concurrency := 5
	information := make(chan int)
	done := make(chan int)

	go func() {
		i := 0
		for {
			information <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go ProcessWorker(information, x)
	}

	<-done

}

func ProcessWorker(information chan int, worker int) {
	for in := range information {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker ", worker, ": ", in)
	}
}
