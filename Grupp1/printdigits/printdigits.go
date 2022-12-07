/*printdigits
Instructions
Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

A line is a sequence of characters preceding the end of line character ('\n').

Usage
$ go run .
0123456789
$
Notions
01-edu/z01*/

package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
