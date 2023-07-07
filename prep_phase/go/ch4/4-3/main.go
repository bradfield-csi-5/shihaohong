// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("before: ", a)
	reverse(&a)
	fmt.Println("after: ", a) // "[5 4 3 2 1 0]"
}

// reverse reverses a slice of ints in place.
// note: this is worse than a slice because the array size is fixed
func reverse(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
