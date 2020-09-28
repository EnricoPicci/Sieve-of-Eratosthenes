// All calls to pNumbers, which is where the prime number is identified and printed, are synchrnous and run in the same goroutine
// the feed of the streams (i.e. writes on channels) passed to the function pNumbers is done concurrently via goroutines
//
// Calculates all prime numbers until a prime number greater than a certain threshold is found.
// The threshold is specified with the parameter "upTo".
// So, if the parameter "upTo" is set to 5, the prime numbers calculated are 2, 3, 5, 7, 11
//
// Note that the program exits when the "main" goroutine ends, which is when a call to "pNumbers" function, which
// is called recursively, exits.
// When "pNumbers" exits, a command is sent to the "generate" function to close the source channel (representing the source
// stream). The closing of the source channel triggers the closing of all the "filteredStream" channels created,
// as can be seen in the "filter" function. But since the various executions of "filter" function calls are all run in their
// own goroutines, there is no guarantee that the "filteredStream" channels are closed before the "main" goroutine exits.
// which anyways terminates the execution of the program.
//
// Run it with the command
// ./prime-numbers-recursive-1 -upTo=100 -printPrime -printCloseChan

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
	quit := make(chan bool)
	go generate(source, quit)
	last := pNumbers(1, source, upTo, quit, printPrime, printCloseChan)
	if printPrime {
		fmt.Println("Last prime number found", last)
	}
}

func generate(sourceStream chan<- int, quit <-chan bool) {
	for i := 2; ; i++ {
		select {
		case <-quit:
			close(sourceStream)
			return
		default:
			sourceStream <- i
		}
	}
}

func pNumbers(i int, inStream <-chan int, upTo int, quit chan<- bool, printPrime bool, printCloseChan bool) int {
	primeNumber := <-inStream // the first number of the stream is a prime number
	if i > upTo {
		quit <- true
		return primeNumber
	}
	if printPrime {
		fmt.Println(i, primeNumber)
	}

	filteredStream := make(chan int)
	// generate a filtered stream out of the inStream
	go filter(inStream, filteredStream, primeNumber, printCloseChan)
	// pass the filtered stream synchrously and recursively to pNumebrs function
	return pNumbers(i+1, filteredStream, upTo, quit, printPrime, printCloseChan)
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
	close(filteredStream)
}
