package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dynamicArray(arrays [][]int64, queries [][]int64) []int64 {
	var lastAnswer int64 = 0
	var answers []int64 = []int64{}

	for i := 0; i < len(queries); i++ {
		query, x, y := queries[i][0], queries[i][1], queries[i][2]
		idx := (x ^ lastAnswer) % int64(len(arrays))

		if query == 1 {
			arrays[idx] = append(arrays[idx], y)
		} else {
			lastAnswer = arrays[idx][y%int64(len(arrays[idx]))]
			answers = append(answers, lastAnswer)
		}
	}

	return answers
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	args := parseInt64Seq(readLine(in))
	n, q := args[0], args[1]

	var arrays [][]int64 = make([][]int64, n)
	var i int64

	for i = 0; i < n; i++ {
		arrays[i] = []int64{}
	}

	var queries [][]int64 = make([][]int64, q)

	for i = 0; i < q; i++ {
		queries[i] = parseInt64Seq(readLine(in))
	}

	var answers []int64 = dynamicArray(arrays, queries)
	for _, v := range answers {
		fmt.Printf("%d\n", v)
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
