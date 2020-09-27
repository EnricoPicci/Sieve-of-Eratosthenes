// primeNumbers returns an Observable representing a stream of prime numbers up to a certain threshold
// Each prime number in the stream is notified to the subscribers of the Observable returned
// since the logic is recursive it eventually hits max stack size
// we are using only Observables and rxjs operators

// Input: a numebr can be passed as input to set the threshold of the prime numbers calculated

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

// 15313 seems the threshold after which we overflow the stack, at least with my current configuration
let _upTo: number;
const upTo = (_upTo = parseInt(process.argv[2])) ? _upTo : 15313;
const pn = primeNumbers(upTo);

pn.subscribe({
    next: (n) => console.log(n),
    error: (e) => console.log('error while generating prime numbers', e),
    complete: () => console.log('All Prime Numbers notified'),
});
