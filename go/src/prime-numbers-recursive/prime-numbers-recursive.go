// the feed of the streams (i.e. writes on channels) passed to the function pNumbers is done concurrently via goroutines
// the identification of the prime number done by the function pNumbers is done synchronously in the same goroutine as main

package main

import (
	"flag"
	"fmt"
)

func main() {
	upTo := flag.Int("upTo", 100, "Threshold of prime numbers found")
	printPrime := flag.Bool("printPrime", false, "Print the prime numbers on the console")
	printCloseChan := flag.Bool("printCloseChan", false, "Print when channels are closed")
	flag.Parse()
	Run(*upTo, *printPrime, *printCloseChan)
}

// Run runs the logic
func Run(upTo int, printPrime bool, printCloseChan bool) {
	source := make(chan int)
	go generate(source, upTo)
	// for the first call to pNumbers the second parameter (i.e. the prime number) is not relevant
	pNumbers(source, printPrime, printCloseChan)
}

func generate(sourceStream chan<- int, upTo int) {
	for i := 2; i < upTo; i++ {
		sourceStream <- i
	}
	close(sourceStream) // the stream is closed to trigger the closing of all goroutines running the filter function
}

var i = 0

func pNumbers(inStream <-chan int, printPrime bool, printCloseChan bool) {
	primeNumber := <-inStream // the first number of the stream is a prime number
	if primeNumber == 0 {
		return
	}
	if printPrime {
		i++
		fmt.Println(i, primeNumber)
	}

	filteredStream := make(chan int)
	go filter(inStream, filteredStream, primeNumber, printCloseChan) // generate a filtered stream out of the inStream
	pNumbers(filteredStream, printPrime, printCloseChan)             // pass the filtered stream synchrously and recursively to pNumebrs function
}

func filter(inStream <-chan int, filteredStream chan<- int, prime int, printCloseChan bool) {
	for i := range inStream {
		if i%prime != 0 {
			filteredStream <- i
		}
	}
	if printCloseChan {
		fmt.Printf("Close stream for prime number %v \n", prime)
	}

	close(filteredStream) // closing this stream triggers the termination of the goroutine running the subsequent filter
}
