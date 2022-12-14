/*strlen
Instructions
Write a function that counts the runes of a string and that returns that count.
Expected function
func StrLen(s string) int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
And its output :

$ go run .
12
$*/
/*
count := 0
for range s
count++
return count
*/

/* Vår lösning
package main

func StrLen(s string) int {
	return len([]rune(s))
}*/

//Stefanies lösning
package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	return len([]rune(s))
}
