Multi-threaded TCP Server
===

## Pre-requisite

Ensure you have [Golang](https://go.dev/) installed. You can
install using the [official installation procedure](https://go.dev/learn/).

## Executing

Using the following command start the webserver. The server will start on port `1729`.

```
$ go run main.go
```

From two different terminal windows execute the following cURL command

```
$ curl http://localhost:1729
```

### Output

Once the cURL commands hit, you should see following output in the terminal session
that is running the server. The output shows interleaved execution of the requests.

```
2023/05/08 20:43:37 Listening on :1729
2023/05/08 20:43:46 new client connected
2023/05/08 20:43:46 processing the request...
2023/05/08 20:43:48 new client connected
2023/05/08 20:43:48 processing the request...
2023/05/08 20:43:51 processing complete...
2023/05/08 20:43:53 processing complete...
```
