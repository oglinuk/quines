package main

import "fmt"

// Go is easier since it has [raw string literals](https://go.dev/ref/spec#String_literals),
// the work-around to not being able to use '`' is the same as '"'.
func main() {
	s := `package main

import "fmt"

// Go is easier since it has [raw string literals](https://go.dev/ref/spec#String_literals),
// the work-around to not being able to use '%c' is the same as '"'.

func main() {
	s := %c%s%c
	fmt.Printf(s, rune(96), rune(96), s, rune(96))
}
`
	fmt.Printf(s, rune(96), rune(96), s, rune(96))
}
