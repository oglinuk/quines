package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	mods = []string {
		"fmt.Fprintf(os.Stderr, %chello, world self-modification one!%c%c%c)",
		"fmt.Fprintf(os.Stderr, %chello, world self-modification two!%c%c%c)",
		"fmt.Fprintf(os.Stderr, %chello, world self-modification three!%c%c%c)",
	}
)

func f() {
	fmt.Fprintf(os.Stderr, "hello, world!")
	old, _ := os.Open("main.go")
	defer old.Close()

	me, _ := os.Create("new.go")
	defer me.Close()

	bs := bufio.NewScanner(old)
	for bs.Scan() {
		me.WriteString(fmt.Sprintf("%s\n", bs.Text()))
		if bs.Text() == "func f() {" {
			me.WriteString(fmt.Sprintf("\t%s\n",fmt.Sprintf(mods[rand.Int()%len(mods)],rune(34),rune(92),rune(110),rune(34))))
		}
	}
	os.Rename("new.go", "main.go")
}

func main() {
	rand.Seed(time.Now().Unix())
	s := `package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	mods = []string {
		"fmt.Fprintf(os.Stderr, %%chello, world self-modification one!%%c%%c%%c)",
		"fmt.Fprintf(os.Stderr, %%chello, world self-modification two!%%c%%c%%c)",
		"fmt.Fprintf(os.Stderr, %%chello, world self-modification three!%%c%%c%%c)",
	}
)

func f() {
	fmt.Fprintf(os.Stderr, "hello, world!")
	old, _ := os.Open("main.go")
	defer old.Close()

	me, _ := os.Create("new.go")
	defer me.Close()

	bs := bufio.NewScanner(old)
	for bs.Scan() {
		me.WriteString(fmt.Sprintf("%%s\n", bs.Text()))
		if bs.Text() == "func f() {" {
			me.WriteString(fmt.Sprintf("\t%%s\n",fmt.Sprintf(mods[rand.Int()%%len(mods)],rune(34),rune(92),rune(110),rune(34))))
		}
	}
	os.Rename("new.go", "main.go")
}

func main() {
	rand.Seed(time.Now().Unix())
	s := %c%s%c

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/self-modifying"},
	}
	_ = modInit.Run()

	// self-modify
	f()

	// self-compile
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "main.go"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}
`

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/self-modifying"},
	}
	_ = modInit.Run()

	// self-modify
	f()

	// self-compile
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "main.go"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}
