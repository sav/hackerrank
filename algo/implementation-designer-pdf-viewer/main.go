package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func designerPdfViewer(heights []int64, word string) int64 {
	max := int64(0)
	for _, c := range word {
		h := heights[int(c-'a')]
		if h > max {
			max = h
		}
	}
	return max * int64(len(word))
}

func main() {
	r := bufio.NewReader(os.Stdin)

	t := strings.Split(line(r), " ")
	h := make([]int64, len(t))
	for i, v := range t {
		var err error
		h[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	w := line(r)
	n := designerPdfViewer(h, w)
	fmt.Println(n)
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
