package main

import "fmt"
import "os"
import "math"

func main() {
	var maxlen int
	_, err := fmt.Scanf("%d", &maxlen)

	if err != nil {
		fmt.Println("Bad Input!")
		os.Exit(1)
	}

	maxfactor := uint(math.Pow(10, float64(maxlen))) - 1

	fmt.Println(maxfactor)
}
