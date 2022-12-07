/*reversebits
Instructions
Write a function that takes a byte, reverses it bit by bit (as shown in the example)
and returns the result.

Expected function
func ReverseBits(oct byte) byte {

}
Example:

1 byte

0010 0110 || / 0110 0100
*/

package main

func ReverseBits(b byte) byte {
	var res byte
	for i := 0; i < 8; i++ {
		res = res << 1
		res += b >> i & 1
	}
	return res
}
