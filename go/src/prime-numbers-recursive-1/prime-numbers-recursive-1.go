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
	flag.Parse()
	Run(*upTo, *printPrime)
}

// Run runs the logic
func Run(upTo int, printPrime bool) {
	source := make(chan int)
	quit := make(chan bool)
	go generate(source, quit)
	last := pNumbers(1, source, upTo, quit, printPrime)
	fmt.Println("Last prime number found", last)
}

func generate(sourceStream chan<- int, quit <-chan bool) {
	for i := 2; ; i++ {
		select {
		case <-quit:
			close(sourceStream)
			return
		default:
			sourceStream <- i
			//fmt.Println("Producer sends", i)
		}
	}
}

// var i = 0

func pNumbers(i int, inStream <-chan int, upTo int, quit chan<- bool, printPrime bool) int {
	primeNumber := <-inStream // the first number of the stream is a prime number
	if i > upTo {
		quit <- true
		return primeNumber
	}
	if printPrime {
		fmt.Println(i, primeNumber)
	}

	filteredStream := make(chan int)
	go filter(inStream, filteredStream, primeNumber)             // generate a filtered stream out of the inStream
	return pNumbers(i+1, filteredStream, upTo, quit, printPrime) // pass the filtered stream synchrously and recursively to pNumebrs function
}

func filter(inStream <-chan int, filteredStream chan<- int, prime int) {
	for i := range inStream {
		if i%prime != 0 {
			filteredStream <- i
		}
	}
	// if printCloseChan {
	// 	fmt.Printf("Close stream for prime number %v \n", prime)
	// }

	//close(filteredStream) // closing this stream triggers the termination of the goroutine running the subsequent filter
}
