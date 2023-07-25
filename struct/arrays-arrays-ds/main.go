package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func reverseArray(a []int64) {
	var n int = len(a)
	var space bool = false
	for i := 0; i < n; i++ {
		if space {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[n-i-1])
		space = true
	}
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	parseInt64(readLine(in))
	a := parseInt64Seq(readLine(in))
	reverseArray(a)
}

func parseInt64Seq(s string) []int64 {
	v := strings.Split(strings.TrimSpace(s), " ")
	n := len(v)
	r := make([]int64, n)
	for i := 0; i < int(n); i++ {
		r[i] = parseInt64(v[i])
	}
	return r
}

func parseInt64(s string) int64 {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	checkError(err)
	return n
}

func parseInt(s string) int {
	return int(parseInt64(s))
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
