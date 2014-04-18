package main

import (
    "fmt"
    "math"
)

func is_leap(year int) bool {
    return math.Mod(float64(year), 400) == 0 || math.Mod(float64(year), 100) != 0 && math.Mod(float64(year), 4) == 0
}

func count_days(year, month, day int) int {
    days := 0
    for y := 1900; y < year; y++ {
        if is_leap(y) {
            days += 366
        } else {
            days += 365
        }
    }
    for m := 1; m < month; m++ {
        if m == 2 {
            if is_leap(year) {
                days += 29
            } else {
                days += 28
            }
        } else {
            if m == 4 || m == 6 || m == 9 || m == 11 {
                days += 30
            } else {
                days += 31
            }
        }
    }
    return days + day
}

func main() {
    count := 0
    for year := 1901; year < 2001; year++ {
        for month := 1; month < 13; month++ {
            if math.Mod(float64(count_days(year, month, 1)), 7) == 0 {
                count++
            }
        }
    }
    fmt.Println(count)
}
