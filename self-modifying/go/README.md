# Self-Modifying Quine in Go 

## How to Use

Since Go requires a module, this implementation invokes the `go mod init`
command as part of the quine. All examples using `go run main.go` can be
replaced with a binary by doing the following.

```BASH
go mod init example.com/self-modifying
go build
```

`go run main.go > original.go`

`go run main.go > original.go && ./main`

`go run main.go > original.go && while :; do ./main; sleep 0.3; done`
