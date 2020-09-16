// with this erxample we calculate prime numbers but we eventually hit max stack size
// we have to use rthe share operator to increase the number of primes we can emit before hitting the stack limit
// we are using only Observables and rxjs operators

import { range, Subject, EMPTY, Observable } from 'rxjs';
import { filter, tap, switchMap, first, share } from 'rxjs/operators';

const primeNumbers$ = new Subject();

primeNumbers$.subscribe({
    next: (d) => console.log(d),
});

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
        tap((d) => {
            primeNumbers$.next(d);
        }),
        switchMap((d) => {
            if (d > end) {
                return EMPTY;
            }
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

// after 8821 we get a stack limit warning
// after 32117 we get a "Maximum call stack size exceeded" error
primeNumbers(2, 32130).subscribe({
    error: (e) => console.log('error r1', e),
    complete: () => console.log('Prime Numbers Done'),
});
