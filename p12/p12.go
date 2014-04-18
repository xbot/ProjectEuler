package main

import (
    "fmt"
    "math"
)

func count_divisor(num int64) int64 {
    var count, factor, lastFactor, exponent int64 = 1, 2, 0, 1

    for factor <= num {
        if math.Mod(float64(num), float64(factor)) == 0 {
            if factor == lastFactor {
                exponent++
            } else {
                count *= (exponent + 1)
                exponent = 1
                lastFactor = factor
            }
            num /= factor
        } else {
            factor++
        }
    }

    return count
}

func main() {
    var count, n int64 = 0, 1
    for count < 500 {
        n++
        count = count_divisor(n * (n + 1) / 2)
    }
    fmt.Println(n*(n+1)/2, count)
}
