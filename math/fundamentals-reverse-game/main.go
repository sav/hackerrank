package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * The first sequences of the reverse game are:
 *
 *  N=1  0
 *  N=2  1 0
 *  N=3  2 0 1
 *  N=4  3 0 2 1
 *  N=5  4 0 3 1 2
 *  N=6  5 0 4 1 3 2
 *  N=7  6 0 5 1 4 2 3
 *  N=8  7 0 6 1 5 2 4 3
 *  N=9  8 0 7 1 6 2 5 3 4
 *
 * Then, in terms of the position we'd have:
 *
 *  N=1  0
 *  N=2  1 0
 *  N=3  1 2 0
 *  N=4  1 3 2 0
 *  N=5  1 3 4 2 0
 *  N=6  1 3 5 4 2 0
 *  N=7  1 3 5 6 4 2 0
 *  N=8  1 3 5 7 6 4 2 0
 *  N=9  1 3 5 7 8 6 4 2 0
 *
 * With this at hand it's possible to infer that the left hand side
 * of a sequence is given by the following equation:
 *
 *  k * 2 + 1
 *
 * While the right hand side can be written as:
 *
 *  (n - k - 1) * 2
 */

func reverseGame(n int64, k int64) int64 {
	if k < n/2 {
		return k*2 + 1
	}
	return (n - k - 1) * 2
}

func main() {
	in := bufio.NewReader(os.Stdin)
	t := parseInt64(readLine(in))
	for ; t > 0; t-- {
		v := parseInt64Seq(readLine(in))
		fmt.Printf("%d\n", reverseGame(v[0], v[1]))
	}
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
