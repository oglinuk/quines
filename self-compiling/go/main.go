package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	s := `package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	s := %c%s%c

	os.Remove("go.mod")

	p, _ := exec.LookPath("go")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/self-compiling"},
	}
	_ = modInit.Run()

	build := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "-o", "self"},
	}
	_ = build.Run()

	fmt.Printf(s, rune(96), s, rune(96))
}
`
	os.Remove("go.mod")

	p, _ := exec.LookPath("go")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/self-compiling"},
	}
	_ = modInit.Run()

	build := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "-o", "self"},
	}
	_ = build.Run()

	fmt.Printf(s, rune(96), s, rune(96))
}
