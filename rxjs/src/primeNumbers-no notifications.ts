// pNumbers returns an Observable that, when subscribed, will print all prime numbers up to a certain threshold
// the print of prime numbers is a side effect and not the result of a "next" function defined by a subscriber
// actually the Observable returned by pNumbers does not notify any value, it just errors at the end
// because the "first" operator is applied to an Observable that just completes before emitting any value

import { range, Observable } from 'rxjs';
import { filter, first, concatMap } from 'rxjs/operators';

const upTo = 100;
const source = range(2, upTo - 2);
function primeNumbersObsWithoutNotifications(source: Observable<number>): Observable<number> {
    return source.pipe(
        first(),
        concatMap((primeNumber) => {
            console.log(primeNumber);
            return primeNumbersObsWithoutNotifications(
                source.pipe(
                    filter((n) => {
                        return n % primeNumber !== 0;
                    }),
                ),
            );
        }),
    );
}

primeNumbersObsWithoutNotifications(source).subscribe({
    next: (d) => console.log('Nottification', d),
    error: (err) => console.error(err),
    complete: () => console.log('DONE'),
});
