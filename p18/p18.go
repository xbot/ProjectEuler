package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    f, err := os.Open("data_p67.txt")
    if nil != err {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    matrix := make([][]int, 0)
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadString('\n')
        if nil != err || io.EOF == err {
            break
        }
        row := make([]int, 0)
        numbers := strings.Split(strings.Replace(line, "\n", "", -1), " ")
        for i := range numbers {
            number, _ := strconv.Atoi(numbers[i])
            row = append(row, number)
        }
        matrix = append(matrix, row)
    }

    for y := 0; y < len(matrix); y++ {
        for x := 0; x < len(matrix[y]); x++ {
            if y > 0 {
                greaterParentPathValue := 0
                if x > 0 {
                    greaterParentPathValue = matrix[y-1][x-1]
                }
                if x < len(matrix[y-1]) && matrix[y-1][x] > greaterParentPathValue {
                    greaterParentPathValue = matrix[y-1][x]
                }
                matrix[y][x] += greaterParentPathValue
            }
        }
    }

    sort.Sort(sort.Reverse(sort.IntSlice(matrix[len(matrix)-1])))
    fmt.Println(matrix[len(matrix)-1][0])
}
