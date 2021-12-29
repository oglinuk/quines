# Self-Compiling Quine in Go 

## How to Use

Since Go requires a module, this implementation invokes the `go mod init`
command as part of the quine.

`go run main.go`

or

```BASH
go mod init example.com/self-compiling
go build
./self-compiling
```
