package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func cavityMap(s []string) []string {
	r := make([]string, len(s))
	copy(r, s)
	h, w := len(s), len(s[0])
	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			x := s[i][j]
			if x > s[i-1][j] && x > s[i+1][j] && x > s[i][j-1] && x > s[i][j+1] {
				t := []byte(r[i])
				t[j] = 'X'
				r[i] = string(t)
			}
		}
	}
	return r
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	n := integer(line(r))
	s := make([]string, 0, n)
	for ; n > 0; n-- {
		s = append(s, line(r))
	}
	t := cavityMap(s)
	for _, v := range t {
		fmt.Println(v)
	}
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
