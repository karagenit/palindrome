package main

import "fmt"
import "os"
import "math"
import "strconv"
import "time"

func isPalin(num int) bool {
	str := strconv.Itoa(num)
	len := len(str)

	for i := 0; i < len; i++ {
		if str[i] != str[len-i-1] { //TODO don't compare all
			return false
		}
	}

	return true
}

func calcPalin(maxfactor int) {
	maxpalin := 0

	for a := maxfactor; a > 0; a-- {
		if (a * a) < maxpalin {
			break
		}
		for b := a; b > 0; b-- {
			if (a * b) < maxpalin {
				break
			}
			if isPalin(a * b) {
				maxpalin = a * b
			}
		}
	}

	fmt.Print("Largest: ")
	fmt.Println(maxpalin)
}

func main() {
	var maxlen int
	_, err := fmt.Scanf("%d", &maxlen)

	if err != nil {
		fmt.Println("Bad Input!")
		os.Exit(1)
	}

	maxfactor := int(math.Pow(10, float64(maxlen))) - 1

	start := time.Now().UnixNano()

	calcPalin(maxfactor)

	time := time.Now().UnixNano() - start

	fmt.Printf("Time: %.2fs\n", float64(time)/1000000000.0)
}
