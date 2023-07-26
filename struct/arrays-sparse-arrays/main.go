package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countMatches(strings []string, match string) int64 {
	var count int64 = 0
	for i := range strings {
		if strings[i] == match {
			count += 1
		}
	}
	return count
}

func sparseArrays(strings []string, queries []string) []int64 {
	var nQueries int64 = int64(len(queries))
	var matches []int64 = make([]int64, nQueries)

	var i int64
	for i = 0; i < nQueries; i++ {
		matches[i] = countMatches(strings, queries[i])
	}

	return matches
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var i int64
	var n int64 = parseInt64(readLine(in))
	var strings []string = make([]string, n)

	for i = 0; i < n; i++ {
		strings[i] = readLine(in)
	}

	n = parseInt64(readLine(in))
	var queries []string = make([]string, n)

	for i = 0; i < n; i++ {
		queries[i] = readLine(in)
	}

	var matches []int64 = sparseArrays(strings, queries)
	for i := range matches {
		fmt.Println(matches[i])
	}
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
