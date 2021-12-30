# Frankenquine in Go

## How to Use

Since Go requires a module, this implementation invokes the `go mod init`
command as part of the quine. All examples using `go run main.go` can be
replaced with a binary by doing the following. This frankenquine produces
a simple hello world program called `monster.go`.

```BASH
go mod init example.com/frankenquine
go build
```

`go run main.go && go run monster.go`
