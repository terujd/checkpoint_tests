/*inter
Instructions
Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

The display will be followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$*/

//Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(str string, character rune) bool {
	for _, r := range str {
		if r == character {
			return false
		}
	}
	return true
}
func main() {
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		result := ""
		for _, r := range str1 {
			for _, r2 := range str2 {
				if r == r2 && check(result, r) == true {
					result = result + string(r)
				}
			}
		}
		for _, r := range result {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
