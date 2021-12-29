package main

import (
	"fmt"
	"os"
)

func f() {
	fmt.Fprintf(os.Stderr, "hello, world!\n")
}

func main() {
	s := `package main

import (
	"fmt"
	"os"
)

func f() {
	fmt.Fprintf(os.Stderr, "hello, world!\n")
}

func main() {
	s := %c%s%c

	f()
	fmt.Printf(s, rune(96), s, rune(96))
}`

	f()
	fmt.Printf(s, rune(96), s, rune(96))
}
