package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func strangeGrid(c int64, r int64) int64 {
	var k int64 = 0
	if c%2 == 0 {
		k = 1
	}
	return ((c/2+(c%2))-1)*10 + ((r-1)*2 + k)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	v := parseInt64Seq(readLine(in))
	fmt.Printf("%d\n", strangeGrid(v[0], v[1]))
}

func parseInt64Seq(s string) []int64 {
	v := strings.Split(strings.TrimSpace(s), " ")
	r := []int64{}
	for i := range v {
		n, err := strconv.ParseInt(v[i], 10, 64)
		checkError(err)
		r = append(r, int64(n))
	}
	return r
}

func parseInt64(s string) int64 {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	checkError(err)
	return int64(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
