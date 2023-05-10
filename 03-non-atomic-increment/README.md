Non-Atomic Increment
===

You can try out the same example on [Replit](https://replit.com/~) and here's
the link to this very code [Non-atomic count++](https://replit.com/@arpitbbhayani/Non-Atomic-Count).

You can also choose to run this locally, and here are the steps.

## Pre-requisite

Ensure you have [GCC](https://gcc.gnu.org/) installed. You can
install using the [official installation procedure](https://gcc.gnu.org/install/).

## Executing

Using the `main.c` fail, generate the assmebly language code.

```
$ gcc -S main.c
```

The the generated `main.s` file containing the generated assmelbly language code.

```
$ cat main.s
```

### Output

Once the `cat` command completes the execution we shou see output something like this.

```
	.file	"main.c"
	.text
	.section	.text.startup,"ax",@progbits
	.p2align 4
	.globl	main
	.type	main, @function
main:
.LFB23:
	.cfi_startproc
	movq	count@GOTPCREL(%rip), %rax
	addl	$1, (%rax)
	xorl	%eax, %eax
	ret
	.cfi_endproc
.LFE23:
	.size	main, .-main
	.globl	count
	.bss
	.align 4
	.type	count, @object
	.size	count, 4
count:
	.zero	4
	.ident	"GCC: (GNU) 11.3.0"
	.section	.note.GNU-stack,"",@progbits
```

The assembly output will vary as per the underlying CPU architecture. If two threads try to run the same `count++`
they will be indulged in a race condition in the `addl` instruction. Note: that the x86 instructions are not atomic.
They are composed of micro-operations.

Inside the processor hardware, the `addl	$1, (%rax)` is actually implemented as three separate "micro-op" instructions:

```
    movl (%rdl), %temp (load)
    addl $1, %temp (add)
    movl %temp, (%rdl) (store)
```

Imagine two threads executing this `addl` instruction at the same time (concurrently). Each thread loads the same value
of (`%rdl`) from memory, then adds one to it in their own separate temporary registers, and then write the same value
back to (`%rdl`) in memory. The last write to memory by each thread will overwrite each other with the same value, and
one increment will essentially be lost.

Temporary register is register which the programmer can't access at all. It is temporarily stored inside microprocessor
which is n-bit operand to the instruction set.

## References

- https://cs.brown.edu/courses/csci0300/2021/notes/l16.html
