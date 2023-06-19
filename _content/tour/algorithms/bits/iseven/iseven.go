package main

import (
	"fmt"
	"iseven"
)

func main() {

	fmt.Println(8, ":", iseven.IsEven(8))
	fmt.Println(15, ":", iseven.IsEven(15))
	fmt.Println(4, ":", iseven.IsEven(4))
}


-- go.mod --

module      "play.ground"

replace (

    "iseven"    => ./iseven
)

-- iseven/iseven.go --
package iseven

// IsEven checks is an integer is even.
func IsEven(num int) bool {

	// This is using the Bitwise AND operator to see if the least significant bit (LSB) is 0.
	// Helpful source: https://catonmat.net/low-level-bit-hacks
	// 0 & 1 = 0 (even number)
	// 1 & 1 = 1 (odd number)
	return num&1 == 0
}
