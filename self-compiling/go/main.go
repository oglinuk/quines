package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := `package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	s := %c%s%c

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", fmt.Sprintf("example.com/self%%d", time.Now().Unix())},
	}
	_ = modInit.Run()

	// self-compile
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "main.go"},
	}
	_ = compile.Run()

	fmt.Printf(s, rune(96), s, rune(96))
}`

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", fmt.Sprintf("example.com/self%d", time.Now().Unix())},
	}
	_ = modInit.Run()

	// self-compile
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "build"},
	}
	_ = compile.Run()

	fmt.Printf(s, rune(96), s, rune(96))
}
