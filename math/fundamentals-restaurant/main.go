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
 * The idea behind this algorithm is straightforward.
 * Find the greatest common divisor of l and b:
 *   d = gcd(l, b)
 * Then calculate the number of equal sized squares:
 *   n = l*b / d^2
 */

func min(x int64, y int64) int64 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func gcd(x int64, y int64) int64 {
	d := min(x, y)
	for ; d > 0; d-- {
		if x%d == 0 && y%d == 0 {
			return d
		}
	}
	return 1
}

func restaurant(l int64, b int64) int64 {
	d := gcd(l, b)
	return l * b / (d * d)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	t := parseInt64(readLine(in))
	for ; t > 0; t-- {
		v := parseInt64Seq(readLine(in))
		fmt.Printf("%d\n", restaurant(v[0], v[1]))
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
