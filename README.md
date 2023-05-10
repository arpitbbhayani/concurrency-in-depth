Concurrency in Depth
===

This repository contains examples to help engineers understand concurrency and multi-threading concepts.
These concepts can be difficult to grasp, even for experienced developers, so the examples in this repository
are designed to help build intuition around multi-threaded code.

These videos are part of the ["Concurrency In Depth"](https://www.youtube.com/playlist?list=PLsdq-3Z1EPT3VjDhjMb5yBsgn0wn2-fjp)
series by [Arpit Bhayani](https://arpitbhayani.me/) which can be found on his YouTube channel
[Asli Engineering by Arpit Bhayani](https://www.youtube.com/@AsliEngineering).

## Introduction

As software systems become more complex, concurrency and multi-threading have become increasingly important.
However, many engineers find it difficult to understand concurrent programs and to reason about their behavior. This repository aims to provide clear,
concise examples that demonstrate key concepts in concurrency and multi-threading.

Concurrency refers to the ability of a system to run multiple tasks simultaneously, while multi-threading refers to the ability of a program to use
multiple threads of execution to perform tasks in parallel. Concurrency and multi-threading are important concepts in software engineering because they
allow for efficient use of system resources and can improve the responsiveness of applications.

## Getting Started

To run the examples in this repository, you will need a basic understanding of concurrency and multi-threading concepts.
If you are new to these concepts, we recommend reading some introductory material before diving into the examples.

To run the examples, navigate to the directory for the example you want to run and follow the instructions in the `README.md` file for that directory.

## Topics

This repository covers the following topics:

- [01-fair-threads](https://youtu.be/2PjlaUnrAMQ) demonstrates how to write fair threads and get optimum performance
- [02-multi-threaded-tcp-server](https://youtu.be/f9gUFy-9uCM) demonstrates why TCP based servers are multi-threaded and how to build one from scratch
- [03-non-atomic-count++](https://youtu.be/kBHd7kn_1EU) demonstrates why count++ operation is not atomic by peeking into the assembly level code

Each example contains a `README.md` file that provides instructions for running the example and an explanation of the concepts demonstrated in the example.

### Pipeline

- [ ] Semaphores and Pessimistic Locking in multi-threaded systems
- [ ] Optimistic Locking and achieveing concurrency without locks
- [ ] Atomic Operations vs Locks - Bechmark and Trade-offs 
- [ ] Concurrency is not parallelism
- [ ] Is python really multi-threaded?
- [ ] Python GIL impact on multi-threading
- [ ] Are fibres lightweight threads? How do they work?
- [ ] What are Goroutines and how are they implemented?
- [ ] Golang Concurrency Patterns
