/*displayalrevm
Instructions
Write a program that displays the alphabet in reverse, with even letters in uppercase,
 and odd letters in lowercase, followed by a newline ('\n').

Usage
$ go run . | cat -e
zYxWvUtSrQpOnMlKjIhGfEdCbA$
$*/

/* Vår lösning
package main

import "github.com/01-edu/z01"

func main() {
	printStr("zYxWvUtSrQpOnMlKjIhGfEdCbA")
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
	for i := 25; i >= 0; i-- {
		var lower = 65 //65 = A in ascii
		if k%2 == 0 {
			z01.PrintRune(rune(lower + i))
		}
		var upper = 97 //97 = a in ascii
		if k&2 != 0 {
			z01.PrintRune(rune(upper + i))
		}
		z01.PrintRune('\n')
	}
}
