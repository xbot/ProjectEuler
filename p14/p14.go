package main

import (
    "fmt"
    "math"
    "time"
)

func collatz(n int) int {
    if n == 1 {
        return 1
    }
    if math.Mod(float64(n), 2) == 0 {
        return collatz(n/2) + 1
    } else {
        return collatz(n*3+1) + 1
    }
}

func main() {
    start := time.Now()

    i, iLen, count, max := 1000000, 0, 0, 0
    for i > 1 {
        iLen = collatz(i)
        if iLen > count {
            count = iLen
            max = i
        }
        i--
    }

    end := time.Now()
    fmt.Println(end.Sub(start), max, count)
}
