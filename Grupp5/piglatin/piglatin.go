/*piglatin
Instructions
Write a program that transforms a string passed as argument in its Pig Latin version.

The rules used by Pig Latin are as follows:

If a word begins with a vowel, just add "ay" to the end.
If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
Only the latin vowels will be considered as vowel(aeiou).
If the word has no vowels, the program should print "No vowels".

If the number of arguments is different from one, the program prints nothing.

Usage
$ go run .
$ go run . pig | cat -e
igpay$
$ go run . Is | cat -e
Isay$
$ go run . crunch | cat -e
unchcray$
$ go run . crnch | cat -e
No vowels$
$ go run . something else | cat -e
$*/

//Stefanies lösning
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}
	s := os.Args[1]
	vowels := "aeiouAEIOU"
	// return the index of the first vowel in s
	firstVowelIndex := func(s string) int {
		for i, ch := range s {
			for _, v := range vowels {
				if ch == v {
					return i
				}
			}
		}
		return -1
	}
	vowel := firstVowelIndex(s)
	novowels := "No vowels"
	if vowel == -1 {
		for _, ch := range novowels {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	} else {
		result := string(s[vowel:]) + string(s[:vowel]) + "ay"
		for _, ch := range result {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
