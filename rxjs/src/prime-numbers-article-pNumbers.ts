// primeNumbers returns an Observable that notifies all prime numbers up to a certain threshold
// since the logic is recursive it eventually hits max stack size
// we have to use the share operator to increase the number of primes we can emit before hitting the stack limit
// we are using only Observables and rxjs operators

import { range, Observable } from 'rxjs';
import { filter, first, concatMap } from 'rxjs/operators';

const upTo = 100;
const source = range(2, upTo - 2);
function primeNumbers(source: Observable<number>): Observable<number> {
    return source.pipe(
        first(),
        concatMap((primeNumber) => {
            console.log(primeNumber);
            return primeNumbers(
                source.pipe(
                    filter((n) => {
                        return n % primeNumber !== 0;
                    }),
                ),
            );
        }),
    );
}

primeNumbers(source).subscribe();
