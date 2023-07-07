package main

import "fmt"

func main() {
	s := []string{"storm", "storm", "storm", "earth", "earth", "fire", "heed", "my", "my", "my", "my", "call"}
	fmt.Println(s)
	s = removeDuplicates(s[:])
	fmt.Println(s)
}

func removeDuplicates(x []string) []string {
	for i := 0; i < len(x)-1; i++ {
		numDuplicates := 0
		j := i + 1
		c1 := x[i]
		c2 := x[j]

		// advance until non-duplicate
		if c1 == c2 {
			for c1 == c2 {
				numDuplicates++
				j++
				c2 = x[j]
			}

			k := i + 1
			for j < len(x) {
				x[k] = x[j]
				j++
				k++
			}
		}

		x = x[:len(x)-numDuplicates]
	}

	return x
}
