package main

import (
	"flag"
	"fmt"
)

// Send the sequence 2, 3, 4, ... to channel 'source'.
func generate(source chan<- int) {
	for i := 2; ; i++ {
		source <- i // Send 'i' to channel 'source'.
	}
}

// Copy the values from channel 'source' to channel 'destination',
// removing those divisible by 'prime'.
func filter(source <-chan int, destination chan<- int, prime int) {
	for i := range source {
		if i%prime != 0 {
			destination <- i // Send 'i' to 'destination'.
		}
	}
}

// Run performs the prime sieve algorithm
func Run(upTo int, printPrime bool) {
	source := make(chan int) // Create a new channel.
	go generate(source)      // Launch Generate goroutine.
	for i := 0; i < 1229; i++ {
		prime := <-source
		if printPrime {
			fmt.Println(i+1, prime)
		}
		destination := make(chan int)
		go filter(source, destination, prime) // launch filter in its own gorouting
		source = destination
	}
}

func main() {
	upTo := flag.Int("upTo", 100, "Threshold of prime numbers found")
	printPrime := flag.Bool("printPrime", false, "Print the prime numbers on the console")
	flag.Parse()
	Run(*upTo, *printPrime)
}
