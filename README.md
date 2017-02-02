# Riemann client (Golang)

## Introduction

Go client library for [Riemann](https://github.com/riemann/riemann).

Features:
* Idiomatic concurrency
* Sending events, state updates, queries.
* Feature parity with the reference implementation written in Ruby.

This client is a fork of Goryman, a Riemann go client written by Christopher Gilbert. Thanks to him ! The initial Goryman repository (https://github.com/bigdatadev/goryman) has been deleted. We used @rikatz fork (https://github.com/rikatz/goryman/) to create this repository.

We renamed the package name of the client `riemanngo` instead of `goryman`

## Installation

To install the package for use in your own programs:

```
go get github.com/riemann/riemann-go-client
```

If you're a developer, Riemann uses [Google Protocol Buffers](https://github.com/golang/protobuf), so make sure that's installed and available on your PATH.

```
go get github.com/golang/protobuf/{proto,protoc-gen-go}
```

## Getting Started

First we'll need to import the library:

```go
import (
    "github.com/riemann/riemann-go-client/"
)
```

Next we'll need to establish a new client:

```go
c := riemanngo.NewGorymanClient("localhost:5555")
err := c.Connect()
if err != nil {
    panic(err)
}
```

Don't forget to close the client connection when you're done:

```go
defer c.Close()
```

Just like the Riemann Ruby client, the client sends small events over UDP by default. TCP is used for queries, and large events. There is no acknowledgement of UDP packets, but they are roughly an order of magnitude faster than TCP. We assume both TCP and UDP are listening on the same port.

Sending events is easy ([list of valid event properties](http://riemann.io/concepts.html)):

```go
err = c.SendEvent(&riemanngo.Event{
    Service: "moargore",
    Metric:  100,
    Tags: []string{"nonblocking"},
})
if err != nil {
    panic(err)
}
```

You can also query events:

```go
events, err := c.QueryEvents("host = \"goryman\"")
if err != nil {
    panic(err)
}
```

The Hostname and Time in events will automatically be replaced with the hostname of the server and the current time if none is specified.

## Copyright

See [LICENSE](LICENSE) document
