package main

import (
	"fmt"
	"strconv"
	"time"
)

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func get_perm(digits []int, number int) string {
	perm, length, counter := "", len(digits), factorial(len(digits)-1)
	for i := 0; i < length; i++ {
		if counter >= number {
			char := strconv.Itoa(digits[i])
			digits = append(digits[:i], digits[i+1:]...)
			perm += char + get_perm(digits, number)
			break
		}
		number -= counter
	}
	return perm
}

func main() {
	startTime := time.Now()
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	perm := get_perm(digits, 1000000)
	fmt.Println(perm, time.Now().Sub(startTime))
}
