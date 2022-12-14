/*searchreplace
Instructions
Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

If the number of arguments is different from 3, the program displays nothing.

If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

Usage
$ go run . "hella there" "a" "o"
hello there
$ go run . "hallo thara" "a" "e"
hello there
$ go run . "abcd" "z" "l"
abcd
$ go run . "something" "a" "o" "b" "c"
$*/

//Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 4 {
		return
	}
	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]
	for _, v := range first {
		if string(v) == second { // string(v) will also read the special characters
			v = rune(third[0])
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
