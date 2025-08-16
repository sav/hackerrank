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
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	line(r) // ignore first line
	a := array(line(r))
	n := len(a)
	sort.Ints(a)
	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", a[i])
	}
	if n > 0 {
		fmt.Printf("%d\n", a[n-1])
	}
}

func integer(s string) int {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func array(s string) []int {
	a := strings.Fields(s)
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = integer(v)
	}
	return b
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
