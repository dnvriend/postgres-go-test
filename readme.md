# postgres-go-test
A small study project on [Go programming language](https://golang.org/) with
the [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) router and postgres [lib/pq](https://godoc.org/github.com/lib/pq),
a Go native driver for postgres.

## How to install
With a correctly configured Go toolchain, navigate to your $GOPATH root folder, most often `~/go`, and type:

```
go get -u github.com/dnvriend/postgres-go-test/...
```

This should create a binary in `$GOPATH/bin`

## How to run
Navigate to `~/go/src/github.com/dnvriend/postgres-go-test` directory and type `make run`

## How to build
To manually build a binary, navigate to `~/go/src/github.com/dnvriend/postgres-go-test` directory
nd type `make install`, this should create a binary in `$GOPATH/bin`.

## Making requests
The API is secured using basicAuth:

```
$ http --auth user:pass :8000/api/actors
```

Have fun!
