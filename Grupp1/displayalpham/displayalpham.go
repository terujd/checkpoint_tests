/*displayalpham
Instructions
Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

Usage
$ go run . | cat -e
aBcDeFgHiJkLmNoPqRsTuVwXyZ$
$*/

/* vår lösning
package main

import "github.com/01-edu/z01"

func main() {
	printStr("aBcDeFgHiJkLmNoPqRsTuVwXyZ")
}

func printStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}*/

//Stefanies lösning
package main

import "github.com/01-edu/z01"

func main() {
	for k := 0; k < 26; k++ {
		var lower = 97 //97 = a in ascii
		if k%2 == 0 {
			z01.PrintRune(rune(lower + k))
		}
		var upper = 65 //65 = A in ascii
		if k&2 != 0 {
			z01.PrintRune(rune(upper + k))
		}
		z01.PrintRune('\n')
	}
}
