// go test -run none -bench ".*" ./... -benchmem
// GOMAXPROCS=x go test -run none -bench ".*" ./... -benchmem

package main

import "testing"

// BenchmarkPrimeNumberRecursive benchmarks the implementation of Prime Sieve recursive
func BenchmarkPrimeNumberRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(1000, false, false)
	}
}
