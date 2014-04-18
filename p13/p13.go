package main

import (
    "bufio"
    "fmt"
    "io"
    "math/big"
    "os"
)

func main() {
    f, err := os.Open("data.txt")
    if nil != err {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()
    reader := bufio.NewReader(f)
    sum, tmp := big.NewInt(0), big.NewInt(0)
    for {
        line, err := reader.ReadString('\n')
        if nil != err || io.EOF == err {
            break
        }
        tmp.SetString(line, 10)
        sum = sum.Add(sum, tmp)
    }
    fmt.Println(sum.String()[:10])
}
