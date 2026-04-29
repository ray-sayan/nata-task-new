package main

import (
	"fmt"
)

func Odd(oddChan chan int, n int) {
	for i := 1; i <= n; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}

func Even(evenChan chan int, n int) {
	for i := 2; i <= n; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func Receiver(oddChan, evenChan chan int, done chan bool) {
	for {
		select {
		case val, ok := <-oddChan:
			if ok {
				fmt.Println(val)
			} else {
				oddChan = nil
			}

		case val, ok := <-evenChan:
			if ok {
				fmt.Println(val)
			} else {
				evenChan = nil
			}
		}

		if oddChan == nil && evenChan == nil {
			break
		}
	}
	done <- true
}

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	oddChan := make(chan int)
	evenChan := make(chan int)
	done := make(chan bool)

	go Odd(oddChan, n)
	go Even(evenChan, n)
	go Receiver(oddChan, evenChan, done)

	<-done
}
