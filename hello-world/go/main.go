package main

import (
	"fmt"
	"os"
)

func main() {
	s := `package main

import (
	"fmt"
	"os"
)

func main() {
	s := %c%s%c

	fmt.Fprintf(os.Stderr, "hello, world!\n")
	fmt.Printf(s, rune(96), s, rune(96))
}`

	fmt.Fprintf(os.Stderr, "hello, world!\n")
	fmt.Printf(s, rune(96), s, rune(96))
}
