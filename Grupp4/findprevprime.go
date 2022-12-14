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

//Stefanies lösning
package main

func isPrime(n int) bool { //this function will keep dividing the given number by all numbers less than it and checking if it's possible returning false
	for i := n - 1; i > 1; i-- { //or if it reaches 1 it will return a true
		if n%i == 0 {
			return false
		}
	}
	return true
}
func FindPrevPrime(nb int) int {
	for i := nb; i > 1; i-- { //this loop will basically start with the given number and then just go one number lower, each time checking if it's a prime number.
		if isPrime(i) == true {
			return i
		}
	}
	return 0 //basically will only happen if the given number is == 1 or less than 1
}
