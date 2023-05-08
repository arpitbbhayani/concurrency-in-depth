Writing Fair Threads
===

## Pre-requisite

Ensure you have [Golang](https://go.dev/) installed. You can
install using the [official installation procedure](https://go.dev/learn/).

## Executing

To understand how to write fair threads, we will run 3 
There are 3 implementation of the same 

1. Simple Sequential Implementation
2. Simple Multi-threaded Implementation
3. Fair Thread Implementation


### Simple Sequential Implementation

```
$ go run simple/main.go
```

#### Output

```
checking till 100000000 found 5761455 prime numbers. took 3m50.2847918s
```

## Simple Multi-threaded Implementation

```
$ go run unfair/main.go
```

#### Output

```
thread 0 [3, 10000003) completed in 11.877999s
thread 1 [10000003, 20000003) completed in 20.0564834s
thread 2 [20000003, 30000003) completed in 24.5238745s
thread 3 [30000003, 40000003) completed in 27.7187593s
thread 4 [40000003, 50000003) completed in 30.1776939s
thread 5 [50000003, 60000003) completed in 33.9010287s
thread 6 [60000003, 70000003) completed in 35.015087s
thread 7 [70000003, 80000003) completed in 38.1051431s
thread 8 [80000003, 90000003) completed in 39.6135365s
thread 9 [90000003, 100000000) completed in 40.9809798s
checking till 100000000 found 5761455 prime numbers. took 40.9811315s
```

## Fair Thread Implementation

```
$ go run fair/main.go
```

#### Output

```
thread 4 completed in 37.0292627s
thread 5 completed in 37.0290735s
thread 2 completed in 37.0289955s
thread 1 completed in 37.0294824s
thread 7 completed in 37.0299688s
thread 9 completed in 37.0306008s
thread 3 completed in 37.0302257s
thread 0 completed in 37.0298607s
thread 6 completed in 37.0301441s
thread 8 completed in 37.0297268s
checking till 100000000 found 5761455 prime numbers. took 37.0311356s
```
