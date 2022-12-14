/*ispowerof2
Instructions
Write a program that determines if a given number is a power of 2.

This program must print true if the given number is a power of 2, otherwise it has to print false.

If there is more than one or no argument, the program should print nothing.

When there is only one argument, it will always be a positive valid int.

Usage :
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
$*/

// Stefanies lösning
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var res string
	if len(args) == 1 {
		n, _ := strconv.Atoi(args[0])
		if n > 0 {
			for n%2 == 0 {
				n /= 2
			}
			if n == 1 {
				res = "true"
			}
		}
		if n == 0 || n != 1 {
			res = "false"
		}
	} else {
		os.Exit(0)
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
