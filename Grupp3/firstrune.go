/*firstrune
Instructions
Write a function that returns the first rune of a string.

Expected function
func FirstRune(s string) rune {

}
Usage
Here is a possible program to test your function :

package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
And its output :

$ go run .
HSO
$
Notions
rune-literals*/

package main

import (
	"github.com/01-edu/z01"
)

func //main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	return rune(s[0])
}
