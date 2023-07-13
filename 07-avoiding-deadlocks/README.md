Avoiding Deadlocks
===

You can try out the same example on [Replit](https://replit.com/~) and here's
the link to this very code [Deadlock-free Code through Total Order](https://replit.com/@arpitbbhayani/Writing-Deadlock-free-Code-throughTotal-Order).

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
how they are trying to acquire locks on certain records. Because the total order
has been imposed on the records while acquiring the lock, the program
never hits a deadlock.


Sample output

```
txn A: wants to acquire lock on record: 0
txn A: acquired lock on record: 0
txn A: wants to acquire lock on record: 2
txn A: acquired lock on record: 2
txn C: wants to acquire lock on record: 0
txn B: wants to acquire lock on record: 0
txn D: wants to acquire lock on record: 0
txn E: wants to acquire lock on record: 0
txn F: wants to acquire lock on record: 1
txn F: acquired lock on record: 1
txn F: wants to acquire lock on record: 2
txn A: released lock on record: 2
txn A: released lock on record: 0
txn F: acquired lock on record: 2
txn C: acquired lock on record: 0
txn C: wants to acquire lock on record: 1
txn A: wants to acquire lock on record: 1
txn F: released lock on record: 2
txn F: released lock on record: 1
txn C: acquired lock on record: 1
txn F: wants to acquire lock on record: 0
txn C: released lock on record: 1
txn C: released lock on record: 0
txn A: acquired lock on record: 1
txn A: wants to acquire lock on record: 2
txn A: acquired lock on record: 2
txn B: acquired lock on record: 0
txn B: wants to acquire lock on record: 1
txn C: wants to acquire lock on record: 1
txn A: released lock on record: 2
txn A: released lock on record: 1
txn B: acquired lock on record: 1
...
...
...
```

## Explanation

### Imposing Total Order

Acquiring locks in a deterministic order can help avoid deadlock. Deadlock occurs when two or more threads are waiting for resources held by each other,
resulting in a circular dependency that cannot be resolved. By enforcing a consistent and deterministic order for acquiring locks,
you can prevent such circular dependencies and mitigate the risk of deadlock.

The key principle to avoid deadlock is to define and enforce a total order on the resources (locks) and ensure that threads always acquire the locks in that order.
This way, threads will never be in a situation where they are waiting for a resource that is held by another thread waiting for a resource held by the original thread,
thus breaking the circular dependency.

To achieve this, you need to establish a lock acquisition hierarchy or order based on the resources involved.
All threads must consistently follow this order when acquiring locks to ensure deadlock-free execution.

It's important to note that enforcing a deterministic lock acquisition order alone may not guarantee freedom from
all potential deadlocks. If your system involves complex dependencies or multiple resources, you may need to analyze the
resource allocation and synchronization mechanisms thoroughly to identify and resolve any potential deadlock scenarios.
