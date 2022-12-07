/*displaylastparam
Instructions
Write a program that displays its last argument, if there is one.

Usage
$ go run . hello there
there
$ go run . "hello there" how are you
you
$ go run . "hello there"
hello there
$ go run .
$*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		printStr(os.Args[len(os.Args)-1])
	}
}

func printStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
