// primeNumbers returns an Observable that notifies all prime numbers up to a certain threshold
// since the logic is recursive it eventually hits max stack size
// we are using only Observables and rxjs operators

import { range, Observable, Subscriber, TeardownLogic } from 'rxjs';
import { filter, first, share, concatMap } from 'rxjs/operators';

function primeNumbers(upTo: number) {
    return new Observable<number>(
        (subscriber: Subscriber<number>): TeardownLogic => {
            // we have to use the share operator to increase the number of primes we can emit before hitting the stack limit
            const startingSource = range(2, upTo - 2).pipe(share());

            function pNumbers(source: Observable<number>): Observable<number> {
                return source.pipe(
                    first(),
                    concatMap((primeNumber) => {
                        if (subscriber.closed) {
                            subscription.unsubscribe();
                        }
                        subscriber.next(primeNumber);
                        return pNumbers(
                            source.pipe(
                                filter((n) => {
                                    return n % primeNumber !== 0;
                                }),
                            ),
                        );
                    }),
                );
            }

            const subscription = pNumbers(startingSource).subscribe({
                error: (err) => {
                    err.name === 'EmptyError' ? subscriber.complete() : subscriber.error(err);
                },
                // complete is never reached since we get first the EmptyError,
            });
        },
    );
}

// xxxx seems the threshold after which we overflow the stack
const pn = primeNumbers(15000);

pn.subscribe({
    next: (n) => console.log(n),
    error: (e) => console.log('error while generating prime numbers', e),
    complete: () => console.log('All Prime Numbers notified'),
});
