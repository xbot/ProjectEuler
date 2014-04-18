package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	line, _ := reader.ReadString('\n')
	names := strings.Split(strings.Replace(line, "\"", "", -1), ",")
	sort.Sort(sort.StringSlice(names))

	val := 0
	for i := 0; i < len(names); i++ {
		tmp := 0
		for j := 0; j < len(names[i]); j++ {
			tmp += int(names[i][j]) - 64
		}
		val += tmp * (i + 1)
	}

	fmt.Println(val)
}
