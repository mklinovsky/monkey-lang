package main

import (
	"fmt"
	"os"

	"github.com/mklinovsky/monkey-lang/repl"
)

func main() {
	fmt.Printf("Monkey lang 🙈\n")
	repl.Start(os.Stdin, os.Stdout)
}
