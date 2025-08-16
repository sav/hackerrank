package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func acmTeamMax(a string, b string) int {
	n := len(a)
	c := 0
	for i := 0; i < n; i++ {
		if a[i] == '1' || b[i] == '1' {
			c++
		}
	}
	return c
}

func acmTeam(s []string) []int {
	n := len(s)
	max := 0
	acc := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			c := acmTeamMax(s[i], s[j])
			if c > max {
				max = c
				acc = 1
			} else if c == max {
				acc++
			}
		}
	}
	return []int{max, acc}
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	t := array(line(r))
	n := t[0]
	subs := make([]string, n)
	for i := int64(0); i < n; i++ {
		subs[i] = line(r)
	}
	a := acmTeam(subs)
	fmt.Printf("%d\n%d\n", a[0], a[1])
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
