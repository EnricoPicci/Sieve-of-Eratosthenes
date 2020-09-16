// with this erxample we calculate prime numbers but we eventually hit max stack size
// we have to use rthe share operator to increase the number of primes we can emit before hitting the stack limit
// we are using only Observables and rxjs operators

import { range, Observable, Subject } from 'rxjs';
import { filter, first, share, concatMap } from 'rxjs/operators';

const primeNumbers$ = new Subject();
primeNumbers$.subscribe({
    next: (d) => console.log(d),
    error: (e) => console.log('error primeNumbers$', e),
    complete: () => console.log('Prime Numbers Done'),
});

function primeNumbers(source: Observable<number>): Observable<number> {
    return source.pipe(
        first(),
        concatMap((primeNumber) => {
            primeNumbers$.next(primeNumber);
            return primeNumbers(source.pipe(filter((n) => n % primeNumber !== 0)));
        }),
    );
}

const end = 15496;
const source = range(2, end - 2).pipe(share());
const pn = primeNumbers(source);

pn.subscribe({
    error: (e) => console.log('error r1', e),
    complete: () => console.log('Generation Done'),
});

setTimeout(() => {
    console.log('Done');
}, 10000);

// function primeNumbers1(source: Observable<number>) {
//   return source.pipe(
//     first(),
//     concatMap((primeNumber) =>
//       primeNumbers1(source.pipe(filter((n) => n % primeNumber !== 0)))
//     )
//   );
// }
