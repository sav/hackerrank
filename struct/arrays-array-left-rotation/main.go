package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func rotateLeft(array []int64, d int64) []int64 {
	var n int64 = int64(len(array))
	var c int64 = 0
	var i int64
	var answer []int64 = make([]int64, n)

	d = d % n
	for i = d; i < n; i++ {
		answer[c] = array[i]
		c += 1
	}
	for i = 0; i < d; i++ {
		answer[c] = array[i]
		c += 1
	}

	return answer
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	args := parseInt64Seq(readLine(in))
	array := parseInt64Seq(readLine(in))

	array = rotateLeft(array, args[1])
	printArray(array)
}

func printArray(array []int64) {
	var n int64 = int64(len(array))
	var i int64
	var space bool = false
	for i = 0; i < n; i++ {
		if space {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", array[i])
		space = true
	}
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
