/*rot13
Instructions
Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

The output will be followed by a newline ('\n').

If the number of arguments is different from 1, the program displays nothing.

Usage
$ go run . "abc"
nop
$ go run . "hello there"
uryyb gurer
$ go run . "HELLO, HELP"
URYYB, URYC
$ go run .
$*/

//Stefanies lösning

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		s := os.Args[1]
		var result string

		for i := 0; i < len(s); i++ {
			if s[i] >= 'a' && s[i] <= 'm' {
				result += string(s[i] + 13)
			}
			if s[i] >= 'n' && s[i] <= 'z' {
				result += string(s[i] - 13)
			}
			if s[i] >= 'A' && s[i] <= 'M' {
				result += string(s[i] + 13)
			}
			if s[i] >= 'N' && s[i] <= 'Z' {
				result += string(s[i] - 13)
			}
			if s[i] < 'A' || s[i] > 'Z' && s[i] < 'a' || s[i] > 'z' {
				result += string(s[i])
			}
		}
		for _, v := range result {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
