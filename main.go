package main

import (
	"fmt"
)

func odd(n int, oddChan chan int) {
	for i := 1; i <= n; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}

func even(n int, evenChan chan int) {
	for i := 2; i <= n; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func receiver(oddChan, evenChan chan int, done chan bool) {
	odd, ok1 := <-oddChan
	even, ok2 := <-evenChan

	for ok1 || ok2 {
		if ok1 {
			fmt.Print(odd, " ")
			odd, ok1 = <-oddChan
		}
		if ok2 {
			fmt.Print(even, " ")
			even, ok2 = <-evenChan
		}
	}
	done <- true
}

func main() {
	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	oddChan := make(chan int)  // unbuffered
	evenChan := make(chan int) // unbuffered
	done := make(chan bool)

	go odd(n, oddChan)
	go even(n, evenChan)
	go receiver(oddChan, evenChan, done)

	<-done
}
