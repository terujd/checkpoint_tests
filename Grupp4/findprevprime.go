/*findprevprime
Instructions
Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

If there are no primes inferior to the int passed as parameter the function should return 0.

Expected function
func FindPrevPrime(nb int) int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.FindPrevPrime(5))
	fmt.Println(piscine.FindPrevPrime(4))
}
And its output :

$ go run .
5
3
$*/

package main
