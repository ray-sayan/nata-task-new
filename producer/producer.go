package main

import "fmt"

// Producer A
func ProducerA(data chan int, turnA chan bool, turnB chan bool, start, end int) {
	for i := start; i <= end; i++ {
		<-turnA
		data <- i
		turnA <- true
	}
	turnB <- true
}

// Producer B
func ProducerB(data chan int, turnB chan bool, turnC chan bool, start, end int) {
	<-turnB
	for i := start; i <= end; i++ {
		data <- i
	}
	turnC <- true
}

// Producer C
func ProducerC(data chan int, turnC chan bool, start, end int) {
	<-turnC
	for i := start; i <= end; i++ {
		data <- i
	}
	close(data)
}

// Receiver
func Receiver(data chan int, done chan bool) {
	for val := range data {
		fmt.Println(val)
	}
	done <- true
}

func main() {
	var aStart, aEnd int
	var bStart, bEnd int
	var cStart, cEnd int

	fmt.Println("Enter range for Producer A (start end):")
	fmt.Scan(&aStart, &aEnd)

	fmt.Println("Enter range for Producer B (start end):")
	fmt.Scan(&bStart, &bEnd)

	fmt.Println("Enter range for Producer C (start end):")
	fmt.Scan(&cStart, &cEnd)

	data := make(chan int) // unbuffered

	turnA := make(chan bool, 1)
	turnB := make(chan bool, 1)
	turnC := make(chan bool, 1)

	done := make(chan bool)

	go ProducerA(data, turnA, turnB, aStart, aEnd)
	go ProducerB(data, turnB, turnC, bStart, bEnd)
	go ProducerC(data, turnC, cStart, cEnd)
	go Receiver(data, done)

	// start with A
	turnA <- true

	<-done
}
