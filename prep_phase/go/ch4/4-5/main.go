package main

import "fmt"

func main() {
	s := []string{"storm", "storm", "storm", "earth", "earth", "fire", "heed", "my", "my", "my", "my", "call"}
	fmt.Println(s) // [storm storm storm earth earth fire heed my my my my call]
	s = removeDuplicates(s[:])
	fmt.Println(s) // [storm earth fire heed my call]

	s = []string{"storm", "earth", "fire", "heed", "my", "call"}
	fmt.Println(s) // [storm earth fire heed my call]
	s = removeDuplicates(s[:])
	fmt.Println(s) // [storm earth fire heed my call]
}

func removeDuplicates(x []string) []string {
	for i := 0; i < len(x)-1; i++ {
		j := i + 1
		c1 := x[i]
		c2 := x[j]

		// advance until non-duplicate
		if c1 == c2 {
			numDuplicates := 0
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
			x = x[:len(x)-numDuplicates]
		}
	}

	return x
}
