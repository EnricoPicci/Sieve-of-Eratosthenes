// with this erxample we calculate prime numbers but we eventually hit max stack size
// we have to use rthe share operator to increase the number of primes we can emit before hitting the stack limit
// we are using only Observables and rxjs operators

import { range, Subject, EMPTY, Observable } from 'rxjs';
import { filter, tap, switchMap, share, map } from 'rxjs/operators';

const primeNumbers$ = new Subject();

primeNumbers$.subscribe({
    next: (d: any) => console.log(d.val),
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
        filter((d) => d.first),
        tap((d) => primeNumbers$.next(d)),
        switchMap((d) => {
            if (d.val > end) {
                return EMPTY;
            }
            return _primeNumbers(pStream.pipe(map((d) => d.val)), d.val, end);
        }),
    );
}

function multiples(source: Observable<number>, n: number) {
    return source.pipe(
        filter((i) => {
            return i % n !== 0;
        }),
        map((d, i) => ({ val: d, first: i === 0 })),
    );
}

primeNumbers(2, 8821).subscribe({
    error: (e) => console.log('error r1', e),
    complete: () => console.log('Prime Numbers Done'),
});
