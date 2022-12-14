/*wdmatch
Instructions
Write a program that takes two string and checks whether it is possible to write,
the first string with characters from the second string.
This rewrite must respect the order in which these characters appear in the second string.

If it is possible, the program displays the string followed by a newline ('\n'),
otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . 123 123
123
$ go run . faya fgvvfdxcacpolhyghbreda
faya
$ go run . faya fgvvfdxcacpolhyghbred
$ go run . error rrerrrfiiljdfxjyuifrrvcoojh
$ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
$ go run .
$*/

/* Vår lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func compare(s, s2 string) {
	res := ""
	for _, e := range s {
		for i, e2 := range s2 {
			if e == e2 {
				res += string(e)
				s2 = s2[i+1:]
				break
			}
		}
	}
	if res == s {
		printStr(res)
	}
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	compare(os.Args[1], os.Args[2])
}*/

//Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {

		first := os.Args[1]
		second := os.Args[2]

		var result string

		for _, v := range first {
			for i, j := range second {
				if v == j {
					result += string(v)
					second = second[i+1:]
					break
				}
			}
			if result == first {
				for _, v := range result {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
			}
		}
	}
}
