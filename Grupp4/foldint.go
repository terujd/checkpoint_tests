/*foldint
Instructions
The function should have as parameters a function,
f func(int, int) int a slice of integers, slice []int and an int acc int.
 For each element of the slice, it should apply the arithmetic function,
  save the result and print it.
  The function will be tested with our own functions Add, Sub, and Mul.

Expected function
func FoldInt(f func(int, int) int, a []int, n int) {

}
Usage
Here is a possible program to test your function:

package main

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
And its output :

$ go run .
99
558
87

93
0
93
$*/

package main

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i <= (len(a))-1; i++ {
		n = f(n, a[i])
	}
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	minus := false
	if n < 0 {
		minus = true
		n = n * -1
	}
	nString := ""
	for n > 0 {
		nString = string(n%10+48) + nString
		n = n / 10
	}
	if minus == true {
		nString = "-" + nString
	}
	for _, r := range nString {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
