// It has the same logic as "prime-numbers-recursive-1" with the difference that there is no command sent to close
// the "source" channel and therefore the program terminates with all the channels still open
//
// Run it with the command
// ./prime-numbers-recursive-2 -upTo=100 -printPrime -printCloseChan

package main

import (
	"flag"
	"fmt"
)

func main() {
	upTo := flag.Int("upTo", 100, "Threshold of prime numbers found")
	printPrime := flag.Bool("printPrime", false, "Print the prime numbers on the console")
	flag.Parse()
	Run(*upTo, *printPrime)
}

// Run runs the logic
func Run(upTo int, printPrime bool) {
	source := make(chan int)
	go generate(source)
	last := pNumbers(1, source, upTo, printPrime)
	if printPrime {
		fmt.Println("Last prime number found", last)
	}
}

func generate(sourceStream chan<- int) {
	for i := 2; ; i++ {
		sourceStream <- i
	}
}

func pNumbers(i int, inStream <-chan int, upTo int, printPrime bool) int {
	primeNumber := <-inStream // the first number of the stream is a prime number
	if i > upTo {
		return primeNumber
	}
	if printPrime {
		fmt.Println(i, primeNumber)
	}

	filteredStream := make(chan int)
	// generate a filtered stream out of the inStream
	go filter(inStream, filteredStream, primeNumber)
	// pass the filtered stream synchrously and recursively to pNumebrs function
	return pNumbers(i+1, filteredStream, upTo, printPrime)
}

func filter(inStream <-chan int, filteredStream chan<- int, prime int) {
	for i := range inStream {
		if i%prime != 0 {
			filteredStream <- i
		}
	}
}
