/*tabmult
Instructions
Write a program that displays a number's multiplication table.

The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.
Usage
$ go run . 9
1 x 9 = 9
2 x 9 = 18
3 x 9 = 27
4 x 9 = 36
5 x 9 = 45
6 x 9 = 54
7 x 9 = 63
8 x 9 = 72
9 x 9 = 81
$ go run . 19
1 x 19 = 19
2 x 19 = 38
3 x 19 = 57
4 x 19 = 76
5 x 19 = 95
6 x 19 = 114
7 x 19 = 133
8 x 19 = 152
9 x 19 = 171
$ go run .

$*/

//Stefanies lösning
package main

//toimii
import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	for i := 1; i < 10; i++ {
		res := convertIntToString(i * n)
		z01.PrintRune(rune(i + 48))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		nr := convertIntToString(n)
		for _, r := range nr {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		for _, r := range res {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

//itoa
func convertIntToString(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(n%10+48) + result
		n /= 10
	}
	return result
}
