# The Sieve or Erastothenes

This repo contains the code to implement the "Sieve of Erastothenes" algorithm using RxJs and Go.

The logic of the implementation is described in the article
https://medium.com/better-programming/prime-numbers-as-streams-with-rxjs-and-go-a18b0292fb5e

## RxJs implementation

`prime-numbers.ts` is the typescript implementation.

`primeNumbers-no-notifications.ts` is a simpler implementation which prints prime numbers as a side effect of the calculation
but does not create an Observable that notifies prime numbers, which on the contrary is what `prime-numbers.ts` does.

### Compile and run the code

- Move to the `rxjs` folder with `cd rxjs`
- Compile the typescript code with `npm run tsc`
- Run `node ./dist/prime-numbers.js 100` to execute the program (in this case up search for prime numbers up to 100). If no input parameter is passed, then a default value is used (see source of `prime-numbers.ts`)

## Go implementation

`prime-numbers-daisy-chain-filter.go` implements the original logic as explained in `https://play.golang.org/p/9U22NfrXeq` or [here](https://risticnikola.com/sieve-of-eratosthenes-in-go)

`prime-numbers-recursive.go` implements the recursive logic explained in the article.

`prime-numbers-recursive-1.go` and `prime-numbers-recursive-2.go` implement variations of the recursive logic
of `prime-numbers-recursive.go`. See the comments at the top of the source files for details.

### Compile the code

- Move to `go` folder with `cd go`.
- Build the executable running the command `go build ./src/prime-numbers-recursive/`

### Launch the code

- Launch the executable running the command `./prime-numbers-recursive`
- Some flags can be specified
  - -upTo=xx : prime numbers are searched up to the xx integer
  - -printPrime : if the flag is set, then prime numbers are printed on the console
  - -printCloseChan : if the flag is set, then a message is printed when channels are closed
- `./prime-numbers-recursive -h` prints the available flags
- Example of command: `./prime-numbers-recursive -upTo=20 -printPrime -printCloseChan`

### Benchmark the 2 implementations

It is possible to benchmark the implementations of the prime-sieve launching the command

`go test -run none -bench ".*" ./... -benchmem`

from within the `go` folder
