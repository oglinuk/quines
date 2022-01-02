# Self-Compiling Quine in Go 

## How to Use

Since Go requires a module, this implementation invokes the `go mod init`
command as part of the quine. All examples using `go run main.go` can be
replaced with a binary by doing the following.

> This implementation does not allow redirection to another file within
> the same directory, since the quine only runs `go build`.

```BASH
go mod init example.com/self-compiling
go build
```

`go run main.go`

`while :; do go run main.go && ls -lh; sleep 1; done`

`mkdir test && go run main.go > test/quine.go && cd test && go build quine.go && ./quine`
