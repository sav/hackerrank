package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func kaprekar(x int64) bool {
	k := len(strconv.FormatInt(x, 10))
	s := strconv.FormatInt(x*x, 10)
	n := len(s)
	l, r := integer(s[:n-k]), integer(s[n-k:])
	return l+r == x
}

func kaprekarNumbers(p int64, q int64) []int64 {
	r := make([]int64, 0, q-p)
	for i := p; i <= q; i++ {
		if kaprekar(i) {
			r = append(r, i)
		}
	}
	return r
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	p := integer(line(r))
	q := integer(line(r))
	a := kaprekarNumbers(p, q)
	n := len(a)

	if n == 0 {
		fmt.Println("INVALID RANGE")
		return
	}
	for i, v := range a {
		fmt.Print(v)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func integer(s string) int64 {
	s = strings.TrimLeft(s, "0")
	if len(s) == 0 {
		return 0
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
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
