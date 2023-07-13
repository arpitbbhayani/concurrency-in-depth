Optimistic Locking
===

You can try out the same example on [Replit](https://replit.com/~) and here's
the link to this very code [Optmistic Locking](https://replit.com/@arpitbbhayani/Optimistic-Locking).

You can also choose to run this locally, and here are the steps.

## Pre-requisite

Ensure you have [GCC](https://gcc.gnu.org/) installed. You can
install using the [official installation procedure](https://gcc.gnu.org/install/).

## Executing

Using the `main.c` fail, generate the assmebly language code.

```
$ gcc -Wall -pthread main.c
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
	.section	.rodata.str1.1,"aMS",@progbits,1
.LC0:
	.string	"inc failed"
	.text
	.p2align 4
	.globl	incrementCount
	.type	incrementCount, @function
incrementCount:
.LFB24:
	.cfi_startproc
	subq	$8, %rsp
	.cfi_def_cfa_offset 16
	movq	count@GOTPCREL(%rip), %rdx
	movl	(%rdx), %eax
	leal	1(%rax), %ecx
	lock cmpxchgl	%ecx, (%rdx)
	jne	.L5
.L2:
	xorl	%edi, %edi
	call	pthread_exit@PLT
.L5:
	leaq	.LC0(%rip), %rdi
	call	puts@PLT
	jmp	.L2
	.cfi_endproc
.LFE24:
	.size	incrementCount, .-incrementCount
	.section	.text.startup,"ax",@progbits
	.p2align 4
	.globl	main
	.type	main, @function
main:
.LFB25:
	.cfi_startproc
	pushq	%r12
	.cfi_def_cfa_offset 16
	.cfi_offset 12, -16
	xorl	%ecx, %ecx
	xorl	%esi, %esi
	subq	$32, %rsp
	.cfi_def_cfa_offset 48
	movq	incrementCount@GOTPCREL(%rip), %r12
	movq	%fs:40, %rax
	movq	%rax, 24(%rsp)
	xorl	%eax, %eax
	leaq	8(%rsp), %rdi
	movq	%r12, %rdx
	call	pthread_create@PLT
	xorl	%ecx, %ecx
	xorl	%esi, %esi
	leaq	16(%rsp), %rdi
	movq	%r12, %rdx
	call	pthread_create@PLT
	movq	8(%rsp), %rdi
	xorl	%esi, %esi
	call	pthread_join@PLT
	movq	16(%rsp), %rdi
	xorl	%esi, %esi
	call	pthread_join@PLT
	movq	24(%rsp), %rax
	subq	%fs:40, %rax
	jne	.L9
	addq	$32, %rsp
	.cfi_remember_state
	.cfi_def_cfa_offset 16
	xorl	%eax, %eax
	popq	%r12
	.cfi_def_cfa_offset 8
	ret
.L9:
	.cfi_restore_state
	call	__stack_chk_fail@PLT
	.cfi_endproc
.LFE25:
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

## `cmpxchgl` instruction

`cmpxchgl` is an assembly language instruction that stands for "compare and exchange long". It is commonly used in multi-threaded programming for atomic operations, specifically for compare-and-swap (CAS) operations.
The `cmpxchgl` instruction compares the value in a register with the value stored at a memory location, and if they are equal, it swaps the value in another register with the value at the memory location.

In simpler terms, `cmpxchgl` is used to perform an atomic compare-and-swap operation, ensuring that multiple threads can safely update shared data without conflicts or race conditions.
