package main

import "fmt"
import "os"
import "math"

func isPalin(i uint) bool {
	return false
}

func main() {
	var maxlen int
	_, err := fmt.Scanf("%d", &maxlen)

	if err != nil {
		fmt.Println("Bad Input!")
		os.Exit(1)
	}

	maxfactor := uint(math.Pow(10, float64(maxlen))) - 1

	for a := maxfactor; a > 0; a-- {
		for b := a; b > 0; b-- {
			if isPalin(a * b) {
				fmt.Println(a * b)
				os.Exit(0)
			}
			//test
			fmt.Println(a * b)
		}
		//test
		fmt.Println("---")
	}
}
