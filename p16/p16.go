package main

import (
    "fmt"
    "math/big"
    "strconv"
)

func main() {
    num, base := big.NewInt(1), big.NewInt(2)
    for power := 0; power < 1000; power++ {
        num.Mul(num, base)
    }
    result := 0
    for i := 0; i < len(num.String()); i++ {
        bit, _ := strconv.Atoi(num.String()[i : i+1])
        result += bit
    }
    fmt.Println(num.String(), result)
}
