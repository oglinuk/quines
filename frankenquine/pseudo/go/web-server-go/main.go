package main

/**
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	time.Sleep(time.Second*5)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})
	fmt.Fprintf(os.Stderr, "Web server running on port 9001 ...\n")
	log.Fatal(http.ListenAndServe(":9001", nil))
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
	m, _ := os.Create("monster.go")
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
				if  _, e := c[t]; !e {
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
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	time.Sleep(time.Second*5)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})
	fmt.Fprintf(os.Stderr, "Web server running on port 9001 ...\n")
	log.Fatal(http.ListenAndServe(":9001", nil))
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
	m, _ := os.Create("monster.go")
	defer m.Close()

	c := make(map[string]struct{})

	bs := bufio.NewScanner(me)
	for bs.Scan() {
		t := bs.Text()
		if t == "/**" {
			for {
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

	gop, _ := exec.LookPath("go")

	// generate Go module
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
		Args: []string{ccp, "monster.c", "-o", "cmain"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}`

	p, _ := exec.LookPath("go")

	// generate Go module
	os.Remove("go.mod")
	modInit := &exec.Cmd{
		Path: p,
		Args: []string{p, "mod", "init", "example.com/pseudo-frankenquine"},
	}
	_ = modInit.Run()

	// self-create
	f()

	// compile
	compile := &exec.Cmd{
		Path: p,
		Args: []string{p, "monster.go", "-o", "monster"},
	}
	_ = compile.Run()

	fmt.Printf(s,rune(96),s,rune(96))
}
