package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	p := array(line(r))
	_, k, q := p[0], p[1], p[2]
	a := array(line(r))
	n := int64(len(a))
	k %= n

	for ; q > 0; q-- {
		i := integer(line(r))
		i = (i + n - k) % n
		fmt.Println(a[i])
	}
}

func integer(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func array(s string) []int64 {
	a := strings.Fields(s)
	b := make([]int64, len(a))
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
