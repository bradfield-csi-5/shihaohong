package utils

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x byte) int {
	return int(pc[x])
}

// CountMatchingBits returns the number of matching bits between two bytes
//
// Takes 32 byte array inputs only
func CountMatchingBits(x [32]byte, y [32]byte) int {
	res := 0
	for i, val := range x {
		b1 := val
		b2 := y[i]
		// XOR the two values to get 0s where the bits match
		// NOT the result of the above to get 1 and run PopCount on it
		// example:
		// b1: 			00110010
		// b2:         	01110000
		// ^(b1 ^ b2):	10111101
		res += PopCount(^(b1 ^ b2))

		// uncomment to check values
		// fmt.Printf("b1 \t%08b\nb2 \t%08b\nb1^b2 \t%08b\n", b1, b2, ^(b1 ^ b2))
		// fmt.Println("res: ", res)
	}

	return res
}
