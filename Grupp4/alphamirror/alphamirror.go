/*alphamirror
Instructions
Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

The case of the letter remains unchanged, for example :

'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

The final result will be followed by a newline ('\n').

If the number of arguments is different from 1, the program prints a new line.

Usage
$ go run . "abc"
zyx$
$
$ go run . "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
$
$ go run .
$*/

// Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	} else {
		s := args[0]
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				c = 'z' - c + 'a'
			} else if c >= 'A' && c <= 'Z' {
				c = 'A' - c + 'Z'
			}
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
