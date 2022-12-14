/*lastrune
Instructions
Write a function that returns the last rune of a string.

Expected function
func LastRune(s string) rune {

}
Usage
Here is a possible program to test your function :

package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
And its output :

$ go run .
!!!
$
Notions
rune-literals*/

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("Hello"))
	z01.PrintRune(LastRune("Salut"))
	z01.PrintRune(LastRune("Ola"))
	z01.PrintRune('\n')
}

/*Vår lösning
func LastRune(s string) rune {
	return rune(s[len(s)-1])
}*/

//Stefanies lösning
func LastRune(s string) rune {
	i := []rune(s)
	j := len(i)
	return i[j-1]
}
