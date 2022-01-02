package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func f() {
/*
	"os"
fmt.Fprintf(os.Stderr, "hello, world!\n")
*/
	// Read our own source code
	me, _ := os.Open("main.go")
	defer me.Close()

	// Create our monster
	m, _ := os.Create("monster.go")
	defer m.Close()

	c := make(map[string]struct{})

	bs := bufio.NewScanner(me)
	for bs.Scan() {
		t := bs.Text()
		if _, e := c[t]; !e {
			if t == "package main" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}
			if t == "import (" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}
			if t == "	\"os\"" {
				m.WriteString(t)
				c[t] = struct{}{}
			}

			if t == "	\"fmt\"" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}

			if t == ")" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}

			if t == "func main() {" {
				m.WriteString(t)
				c[t] = struct{}{}
			}

			if _, e := c["func main() {"]; e {
				if t == "fmt.Fprintf(os.Stderr, \"hello, world!\\n\")" {
					m.WriteString(t)
					c[t] = struct{}{}
				}

				if t == "}" {
					m.WriteString(t)
					c[t] = struct{}{}
				}
			}
		}
	}
}

func main() {
	s := `package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func f() {
/*
	"os"
fmt.Fprintf(os.Stderr, "hello, world!\n")
*/
	// Read our own source code
	me, _ := os.Open("main.go")
	defer me.Close()

	// Create our monster
	m, _ := os.Create("monster.go")
	defer m.Close()

	c := make(map[string]struct{})

	bs := bufio.NewScanner(me)
	for bs.Scan() {
		t := bs.Text()
		if _, e := c[t]; !e {
			if t == "package main" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}
			if t == "import (" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}
			if t == "	\"os\"" {
				m.WriteString(t)
				c[t] = struct{}{}
			}

			if t == "	\"fmt\"" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}

			if t == ")" {
				m.WriteString(t+string(rune(10)))
				c[t] = struct{}{}
			}

			if t == "func main() {" {
				m.WriteString(t)
				c[t] = struct{}{}
			}

			if _, e := c["func main() {"]; e {
				if t == "fmt.Fprintf(os.Stderr, \"hello, world!\\n\")" {
					m.WriteString(t)
					c[t] = struct{}{}
				}

				if t == "}" {
					m.WriteString(t)
					c[t] = struct{}{}
				}
			}
		}
	}
}

func main() {
	s := %c%s%c

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/traditional-frankenquine"},
	}
	_ = modInit.Run()

	// self-create
	f()

	// compile monster
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
		Args: []string{p, "mod", "init", "example.com/traditional-frankenquine"},
	}
	_ = modInit.Run()

	// self-create
	f()

	// compile monster
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "build", "monster.go"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}
