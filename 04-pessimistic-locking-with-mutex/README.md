Pessimistic Locking with Mutex
===

## Pre-requisite

Ensure you have [Golang](https://go.dev/) installed. You can
install using the [official installation procedure](https://go.dev/learn/).

## Executing

Using the following command start the webserver. The server will start on port `1729`.

```
$ go run main.go
```

### Output

Once the commands completes the execution, you should see following output in the terminal session.

```
count (without lock) 895405
count (with mutex) 1000000
```
