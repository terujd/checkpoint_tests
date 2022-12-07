/*displayalpham
Instructions
Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

Usage
$ go run . | cat -e
aBcDeFgHiJkLmNoPqRsTuVwXyZ$
$*/

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
}
