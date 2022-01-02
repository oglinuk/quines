package main

/**
#include <stdio.h>

int main()
{
	fprintf(stderr, "hello, world!\n");
	return 0;
}
**/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func f() {
	// Read our own source code
	me, _ := os.Open("main.go")
	defer me.Close()

	// Create our monster
	m, _ := os.Create("monster.c")
	defer m.Close()

	c := make(map[string]struct{})

	bs := bufio.NewScanner(me)
	for bs.Scan() {
		t := bs.Text()
		if t == "/**" {
			for bs.Scan() {
				t = bs.Text()
				if t == "**/" {
					break
				}
				if _, e := c[t]; !e {
					c[t] = struct{}{}
					m.WriteString(t+string(rune(10)))
				}
			}
		}
	}
}

func main() {
	s := `package main

/**
#include <stdio.h>

int main()
{
	fprintf(stderr, "hello, world!\n");
	return 0;
}
**/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func f() {
	// Read our own source code
	me, _ := os.Open("main.go")
	defer me.Close()

	// Create our monster
	m, _ := os.Create("monster.c")
	defer m.Close()

	c := make(map[string]struct{})

	bs := bufio.NewScanner(me)
	for bs.Scan() {
		t := bs.Text()
		if t == "/**" {
			for bs.Scan() {
				t = bs.Text()
				if t == "**/" {
					break
				}
				if _, e := c[t]; !e {
					c[t] = struct{}{}
					m.WriteString(t+string(rune(10)))
				}
			}
		}
	}
}


func main() {
	s := %c%s%c

	// generate Go module
	gop, _ := exec.LookPath("go")
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: gop,
		Args: []string{gop, "mod", "init", "example.com/pseudo-frankenquine"},
	}
	_ = modInit.Run()

	// self-create
	f()

	// compile
	ccp, _ := exec.LookPath("cc")
	compile := &exec.Cmd{
		Path: ccp,
		Args: []string{ccp, "monster.c", "-o", "monster"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}`

	// generate Go module
	gop, _ := exec.LookPath("go")
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: gop,
		Args: []string{gop, "mod", "init", "example.com/pseudo-frankenquine"},
	}
	_ = modInit.Run()

	// self-create
	f()

	// compile
	ccp, _ := exec.LookPath("cc")
	compile := &exec.Cmd{
		Path: ccp,
		Args: []string{ccp, "monster.c", "-o", "monster"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}
