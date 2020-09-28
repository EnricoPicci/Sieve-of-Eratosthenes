// All calls to pNumbers, which is where the prime number is identified and printed, are synchrnous and run in the same goroutine
// the feed of the streams (i.e. writes on channels) passed to the function pNumbers is done concurrently via goroutines
//
// Calculates all prime numbers up to a certain integer specified in the parameter "upTo".
// So, if the parameter "upTo" is set to 10, the prime numbers calculated are 2, 3, 5, 7
//
// Note that the program completes when all the channels (the initial source channel and the following "filteredStream" chanels)
// are closed, as it is explained here https://medium.com/better-programming/prime-numbers-as-streams-with-rxjs-and-go-a18b0292fb5e
//
// Run it with the command
// ./prime-numbers-recursive -upTo=100 -printPrime -printCloseChan

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
	pNumbers(1, source, printPrime, printCloseChan)
	if printCloseChan {
		fmt.Println("end of processing")
	}
}

func generate(sourceStream chan<- int, upTo int) {
	for i := 2; i < upTo; i++ {
		sourceStream <- i
	}
	// the stream is closed to trigger the closing of all goroutines running the filter function
	close(sourceStream)
}

func pNumbers(i int, inStream <-chan int, printPrime bool, printCloseChan bool) {
	primeNumber, more := <-inStream // the first number of the stream is a prime number
	if !more {
		return
	}
	if printPrime {
		i++
		fmt.Println(i, primeNumber)
	}

	filteredStream := make(chan int)
	// generate a filtered stream out of the inStream
	go filter(inStream, filteredStream, primeNumber, printCloseChan)
	// pass the filtered stream synchrously and recursively to pNumebrs function
	pNumbers(i+1, filteredStream, printPrime, printCloseChan)
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
	// closing this stream triggers the termination of the goroutine running the subsequent filter
	// The filteredStream where all the non-multiple of 3 are written is the inStream used to build
	// the subsequent stream, i.e. the filteredStream where all the non-multiple of 5 are written
	close(filteredStream)
}
