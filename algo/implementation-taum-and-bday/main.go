package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	if wc+z < bc {
		bc = wc + z
	} else if bc+z < wc {
		wc = bc + z
	}
	return b*bc + w*wc
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	n := integer(line(r))
	for ; n > 0; n-- {
		amounts := array(line(r))
		b, w := amounts[0], amounts[1]
		prices := array(line(r))
		bc, wc, z := prices[0], prices[1], prices[2]
		answer := taumBday(b, w, bc, wc, z)
		fmt.Println(answer)
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
