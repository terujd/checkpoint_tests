/*hello
Instructions
Write a program that displays "Hello World!" followed by a newline ('\n').

Usage
$ go run .
Hello World!
$*/

/* Vår lösning
package main

import "github.com/01-edu/z01"

func main() {
	printStr("Hello World!")
}

func printStr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}*/

// Stefanies lösning
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := "Helllo World!"
	for _, v := range a {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
