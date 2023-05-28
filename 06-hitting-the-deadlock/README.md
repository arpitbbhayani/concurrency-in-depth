Hitting the Deadlock
===

You can try out the same example on [Replit](https://replit.com/~) and here's
the link to this very code [Non-atomic count++](https://replit.com/@arpitbbhayani/Non-Atomic-Count).

You can also choose to run this locally, and here are the steps.

## Pre-requisite

Ensure you have [GCC](https://gcc.gnu.org/) installed. You can
install using the [official installation procedure](https://gcc.gnu.org/install/).

## Executing

Compile and generate binary from `main.c` file.

```
$ gcc main.c -lpthread -Wall
```

### Output

Run the executable `a.out` and you would output containing the transaction and
how they are trying to acquire locks on certain records. After few iterations
you would see the execution hitting deadlock and hanging.

Note: the number of iterations after which the program hangs totally depends
on how CPU has scheduled the threads.

Sample output

```
txn A: wants to acquire lock on record: 1
txn A: acquired lock on record: 1
txn A: wants to acquire lock on record: 0
txn A: acquired lock on record: 0
txn B: wants to acquire lock on record: 1
txn C: wants to acquire lock on record: 2
txn C: acquired lock on record: 2
txn C: wants to acquire lock on record: 1
txn D: wants to acquire lock on record: 1
txn E: wants to acquire lock on record: 2
txn F: wants to acquire lock on record: 1
txn A: released lock on record: 0
txn A: released lock on record: 1
txn B: acquired lock on record: 1
txn B: wants to acquire lock on record: 0
txn B: acquired lock on record: 0
txn A: wants to acquire lock on record: 2
txn B: released lock on record: 0
txn B: released lock on record: 1
txn C: acquired lock on record: 1
txn B: wants to acquire lock on record: 0
txn B: acquired lock on record: 0
txn B: wants to acquire lock on record: 2
txn C: released lock on record: 1
txn C: released lock on record: 2
txn D: acquired lock on record: 1
txn D: wants to acquire lock on record: 2
txn E: acquired lock on record: 2
txn E: wants to acquire lock on record: 0
txn C: wants to acquire lock on record: 0
```
