package main

import "fmt"
import "os"
import "math"
import "strconv"

func isPalin(num int) bool {
	str := strconv.Itoa(num)
	len := len(str)

	for i := 0; i < len; i++ {
		if str[i] != str[len-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var maxlen int
	_, err := fmt.Scanf("%d", &maxlen)

	if err != nil {
		fmt.Println("Bad Input!")
		os.Exit(1)
	}

	maxfactor := int(math.Pow(10, float64(maxlen))) - 1
	maxpalin := 0

	for a := maxfactor; a > 0; a-- {
		if (a * a) < maxpalin {
			fmt.Println("Ye")
			fmt.Println(maxpalin)
			os.Exit(0)
		}
		for b := a; b > 0; b-- {
			if (a*b) > maxpalin && isPalin(a*b) { //which order is best?
				maxpalin = a * b
			}
			//test
			fmt.Println(a * b)
		}
		//test
		fmt.Println("---")
	}
}
