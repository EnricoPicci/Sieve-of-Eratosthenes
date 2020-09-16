// with this erxample we calculate prime numbers but we eventually hit max stack size
// we have to use rthe share operator to increase the number of primes we can emit before hitting the stack limit
// we are using only Observables and rxjs operators

import { range, EMPTY, Observable } from 'rxjs';
import { filter, switchMap, first, share } from 'rxjs/operators';

// const primeNumbers$ = new Subject();

// primeNumbers$.subscribe({
//   next: (d) => console.log(d),
// });

let lastPrimeNumber = 0;

export function primeNumbers(start: number, end: number) {
    const source = range(start, end - start);
    return _primeNumbers(source, start, end);
}
export function _primeNumbers(source: Observable<number>, n: number, end: number) {
    const pStream = multiples(source, n).pipe(
        // share here is critical to avoid hitting stack limits too early
        share(),
    );
    return pStream.pipe(
        first(),
        // tap((d) => {
        //   setTimeout(() => {
        //     primeNumbers$.next(d);
        //   }, 0);
        // }),
        switchMap((d) => {
            if (d > end) {
                return EMPTY;
            }
            lastPrimeNumber = d;
            return _primeNumbers(pStream, d, end);
        }),
    );
}

function multiples(source: Observable<number>, n: number) {
    return source.pipe(
        filter((i) => {
            return i % n !== 0;
        }),
    );
}

// after 40487 we get a "Maximum call stack size exceeded" error
const pn = primeNumbers(2, 5487);
pn.subscribe({
    error: (e) => console.log('error r1', e),
    complete: () => console.log('Prime Numbers Done'),
});

setTimeout(() => {
    console.log('Last Prime Number found', lastPrimeNumber);
}, 10000);
