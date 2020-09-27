# The Sieve or Erastothenes

This repo contains the code to implement the "Sieve of Erastothenes" algorithm using RxJs and Go.

The logic of the implementation is described in the article
https://github.com/EnricoPicci/observable-fs

## RxJs implementation

`prime-numbers.ts` is the typescript implementation.

`primeNumbers-no-notifications.ts` is a simpler implementation which prints prime numbers as a side effect of the calculation
but does not create an Observable that notifies prime numbers, which on the contrary is what `prime-numbers.ts` does.

### Compile and run the code

- Move to the `rxjs` folder with `cd rxjs`
- Compile the typescript code with `npm run tsc`
- Run `node ./dist/prime-numbers.js 100` to execute the program (in this case up search for prime numbers up to 100). If no input parameter is passed, then a default value is used (see source of `prime-numbers.ts`)

## Go implementation

`prime-numbers-recursive.go` implements the recursive logic explained in the article.

`prime-numbers-daisy-chain-filter.go` implements the original logic as explained in `https://play.golang.org/p/9U22NfrXeq` or [here](https://risticnikola.com/sieve-of-eratosthenes-in-go)

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
- `./prime-numbers-recursive -upTo=20 -printPrime -printCloseChan`

### Benchmark the 2 implementations

It is possible to benchmark the 2 implementations, the recursive one (`prime-numbers-recursive.go`)
and the daisy-chain one (`prime-numbers-daisy-chain-filter.go`)
