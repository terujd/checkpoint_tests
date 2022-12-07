/*lastword
Instructions
Write a program that takes a string and displays its last word, followed by a newline ('\n').

A word is a section of string delimited by spaces or by the start/end of the string.

The output will be followed by a newline ('\n').

If the number of arguments is different from 1,
or if there are no word, the program displays nothing.

Usage
$ go run . "FOR PONY" | cat -e
PONY$
$ go run . "this        ...       is sparta, then again, maybe    not" | cat -e
not$
$ go run . "  "
$ go run . "a" "b"
$ go run . "  lorem,ipsum  " | cat -e
lorem,ipsum$
$*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		printStr(string(len(os.Args)))
	}
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
