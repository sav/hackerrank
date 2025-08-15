package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func pickingNumbers(a []int64) int64 {
	var b []int64 = make([]int64, 101)
	for _, v := range a {
		b[v]++
	}
	max := int64(0)
	for i := 1; i < 101; i++ {
		if b[i-1]+b[i] > max {
			max = b[i-1] + b[i]
		}
	}
	return max
}

func main() {
	r := bufio.NewReader(os.Stdin)
	line(r) // ignore the first line (array size)
	s := line(r)
	a := strings.Split(s, " ")
	n := len(a)
	b := make([]int64, n)

	for i, v := range a {
		var err error
		b[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	c := pickingNumbers(b)
	fmt.Println(c)
}

func line(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	if err == io.EOF {
		panic(err)
	}
	s := strings.TrimRight(string(line), "\r\n")
	s = strings.TrimSpace(s)
	return s
}
