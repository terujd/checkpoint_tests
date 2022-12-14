/*chunk
Instructions
Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

If the size is 0 it should print a newline ('\n').
Expected function
func Chunk(slice []int, size int) {

}
Usage
Here is a possible program to test your function :

package main

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
And its output :

$ go run .
[]

[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
$*/

//Stefanies lösning
package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println("")
	} else if len(slice) == 0 {
		fmt.Println(slice)
	} else {
		a := make([][]int, 0, size)
		for size < len(slice) {
			a = append(a, slice[0:size])
			slice = slice[size:]
		}
		fmt.Println(append(a, slice))
	}
}
