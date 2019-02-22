package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gotool: ")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "stdin", os.Stdin, parser.ImportsOnly)
	if err != nil {
		log.Fatal(err)
	}

	for _, imp := range f.Imports {
		name, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			log.Fatal(errors.Wrap(err, fmt.Sprintf("cannot unquote value at %d", imp.Path.Pos())))
		}

		log.Printf("BEG go install %s\n", name)
		cmd := exec.Command("go", "install", name)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		log.Printf("END go install %s\n", name)
	}
}
