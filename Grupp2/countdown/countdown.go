/*countdown
Instructions
Write a program that displays all digits in descending order, followed by a newline ('\n').

Usage
$ go run .
9876543210
$*/

package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
