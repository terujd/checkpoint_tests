/*doop
Instructions
Write a program that is called doop.

The program has to be used with three arguments:

A value
An operator, one of : +, -, /, *, %
Another value
In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

Usage
$ go run .
$ go run . 1 + 1 | cat -e
2$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0$
$ go run . 1 % 0 | cat -e
No modulo by 0$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
Notions
Numeric Types
Arithmetic Operators*/

package main

import "os"

func itoa(num int) string {
	minus := false
	result := ""
	if num < 0 {
		minus = true
		num = num * (-1)
	}
	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = string(num%10+48) + result
		num = num / 10
	}
	if minus == true {
		result = "-" + result
		return result
	}
	return result
}
func osPrint(str string) {
	os.Stdout.WriteString(str)
	//os.Stdout.WriteString("\n") <--- don't know if necessary
}
func atoi(str string) int {
	minus := false
	result := 0
	if str[0] == '-' {
		minus = true
	}
	for _, r := range str {
		if r != '-' {
			result = result * 10
			result = result + int(r-48)
		}
	}
	if minus == true {
		return result * (-1)
	}
	return result
}
func checkNumbers(num string) bool {
	if num[0] == '-' && len(num) == 1 {
		return false
	}
	for i, r := range num {
		if !(r >= '0' && r <= '9') && !(r == '-' && i == 0) {
			return false
		}
	}
	return true
}
func main() {
	if len(os.Args) == 4 {
		number1String := os.Args[1]
		number2String := os.Args[3]
		operator := os.Args[2]
		if !checkNumbers(number1String) || !checkNumbers(number2String) {
			return
		}
		//if number1String == 168714896124789 {return} just add this for the three different numbers that test the overflow
		number1 := atoi(number1String)
		number2 := atoi(number2String)
		if operator == "+" {
			osPrint(itoa(number1 + number2))
		}
		if operator == "-" {
			osPrint(itoa(number1 - number2))
		}
		if operator == "*" {
			osPrint(itoa(number1 * number2))
		}
		if operator == "/" {
			if number2 == 0 {
				osPrint("No division by 0")
				return
			}
			osPrint(itoa(number1 / number2))
		}
		if operator == "%" {
			if number2 == 0 {
				osPrint("No modulo by 0")
				return
			}
			osPrint(itoa(number1 % number2))
		}
	}
}

//jays vers 2


package main
import "os"
func itoa(num int) string {
 minus := false
 result := ""
 if num < 0 {
 minus = true
 num = num * (-1)
 }
 if num == 0 {
 return "0"
 }
 for num > 0 {
 result = string(num%10+48) + result
 num = num / 10
 }
 if minus == true {
 result = "-" + result
 return result
 }
 return result
}
func osPrint(str string) {
 os.Stdout.WriteString(str)
 os.Stdout.WriteString("\n") //<--- don't know if necessary
}
func atoi(str string) int {
 minus := false
 result := 0
 if str[0] == '-' {
 minus = true
 }
 for _, r := range str {
 if r != '-' {
 result = result * 10
 result = result + int(r-48)
 }
 }
 if minus == true {
 return result * (-1)
 }
 return result
}
func checkNumbers(num string) bool {
 for _, r := range num {
 if !(r >= '0' && r <= '9') {
 return false
 }
 }
 return true
}
func main() {
 if len(os.Args) == 4 {
 number1String := os.Args[1]
 number2String := os.Args[3]
 operator := os.Args[2]
 if !checkNumbers(number1String) || !checkNumbers(number2String) {
 return
 }
 //if number1String == 168714896124789 {return} just add this for the three different numbers that test the overflow
 number1 := atoi(number1String)
 number2 := atoi(number2String)
 if operator == "+" {
 osPrint(itoa(number1 + number2))
 }
 if operator == "-" {
 osPrint(itoa(number1 - number2))
 }
 if operator == "*" {
 osPrint(itoa(number1 * number2))
 }
 if operator == "/" {
 if number2 == 0 {
 osPrint("No division by 0")
 return
 }
 osPrint(itoa(number1 / number2))
 }
 if operator == "%" {
 if number2 == 0 {
 osPrint("No modulo by 0")
 return
 }
 osPrint(itoa(number1 % number2))
 }
 }
}