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
	r := bufio.NewReader(os.Stdin)
	t, err := strconv.ParseInt(line(r), 10, 64)
	if err != nil {
		panic(err)
	}
	for ; t > 0; t-- {
		s := line(r)
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		c := 0
		for _, r := range s {
			m := int64(r - '0')
			if m > 0 && n%m == 0 {
				c++
			}
		}
		fmt.Println(c)
	}
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
