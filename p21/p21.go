package main

import (
	"fmt"
	"math"
	"time"
)

func sum_proper_factors(n int) int {
	sum, sqrt := 1, math.Sqrt(float64(n))

	start, step := 2, 1
	if n%2 == 1 {
		start, step = 3, 2
	}
	for i := start; i <= int(sqrt); i += step {
		if n%i == 0 {
			sum += i + n/i
		}
	}

	if sqrt == float64(int(sqrt)) {
		sum -= int(sqrt)
	}

	return sum
}

func main() {
	result, startTime := 0, time.Now()

	for i := 1; i < 10000; i++ {
		iSum := sum_proper_factors(i)
		if iSum > i {
			if i == sum_proper_factors(iSum) {
				result += i + iSum
			}
		}
	}

	fmt.Println(result, time.Now().Sub(startTime))
}
