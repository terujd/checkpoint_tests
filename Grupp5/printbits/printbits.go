/*printbits
Instructions
Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

If the the argument is not a number, the program should print 00000000.
Usage :
$ go run . 1
00000001$
$ go run . 192
11000000$
$ go run . a
00000000$
$ go run . 1 1
$ go run .
$*/

package main

import "github.com/01-edu/z01"

func PrintBits(octet byte) {
	for i := 7; i >= 0; i-- {
		if octet&(1<<uint(i)) != 0 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
}
