/*union
Instructions
Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').

If the number of arguments is different from 2, then the program displays a newline ('\n').

Usage
$ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
$
$ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
$
$ go run . rien "cette phrase ne cache rien" | cat -e
rienct phas$
$
$ go run . | cat -e
$
$ go run . rien | cat -e
$
*/

//Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(result string, character rune) bool { //checks if a character already exists in result and returns either true or false
	for _, r := range result {
		if r == character {
			return false
		}
	}
	return true
}
func main() {
	if len(os.Args) == 3 {
		str := os.Args[1] + os.Args[2]
		result := ""
		for _, r := range str { // loops through combined strings and saves only the unique characters by running each character in the check function
			if check(result, r) == true {
				result = result + string(r)
			}
		}
		for _, r := range result {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
