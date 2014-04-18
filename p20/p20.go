package main

import (
    "fmt"
    "math/big"
    "strconv"
)

func main() {
    product := big.NewInt(1)
    for i := 1; i < 101; i++ {
        product.Mul(product, big.NewInt(int64(i)))
    }
    sum := 0
    for idx := range product.String() {
        tmp, _ := strconv.Atoi(product.String()[idx : idx+1])
        sum += tmp
    }
    fmt.Println(sum)
}
