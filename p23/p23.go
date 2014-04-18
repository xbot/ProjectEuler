package main

import (
	"fmt"
	"math"
	"time"
)

func sum_proper_factors(n int) int {
	result, sqrt := 1, math.Sqrt(float64(n))

	start, step := 2, 1
	if n%2 == 1 {
		start, step = 3, 2
	}
	for i := start; i <= int(sqrt); i += step {
		if n%i == 0 {
			result += i + n/i
		}
	}

	if sqrt == float64(int(sqrt)) {
		result -= int(sqrt)
	}

	return result
}

func main() {
	result, limit, abundants, startTime := 0, 28124, make(map[int]bool), time.Now()

	for n := 1; n < limit; n++ {
		if sum_proper_factors(n) > n {
			abundants[n] = true
		}
		isSumOfTwoAbundants := false
		for k := range abundants {
			if abundants[n-k] == true {
				isSumOfTwoAbundants = true
				break
			}
		}
		if !isSumOfTwoAbundants {
			result += n
		}
	}

	fmt.Println(result, time.Now().Sub(startTime))
}
