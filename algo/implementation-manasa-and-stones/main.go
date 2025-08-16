package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func stones(n int64, a int64, b int64) []int64 {
	r := make([]int64, 0, n)
	if a > b {
		t := b
		b = a
		a = t
	}
	c := (n - 1) * a
	r = append(r, c)
	for i := int64(0); i < n-1; i++ {
		c = c - a + b
		if r[len(r)-1] != c {
			r = append(r, c)
		}
	}
	return r
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	t := integer(line(r))
	for ; t > 0; t-- {
		n := integer(line(r))
		a := integer(line(r))
		b := integer(line(r))
		v := stones(n, a, b)
		parray(v)
	}
}

func parray(v []int64) {
	n := len(v)
	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Printf("%d\n", v[n-1])
}

func integer(s string) int64 {
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
