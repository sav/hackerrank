package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func twoDimArray(a [][]int64) int64 {
	r := int64(-99)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := a[i][j] + a[i][j+1] + a[i][j+2] + a[i+1][j+1] + a[i+2][j] + a[i+2][j+1] + a[i+2][j+2]
			if sum > r {
				r = sum
			}
		}
	}
	return r
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	var a [][]int64 = make([][]int64, 6)
	for i := 0; i < 6; i++ {
		a[i] = parseInt64Seq(readLine(in))
	}
	fmt.Println(twoDimArray(a))
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
