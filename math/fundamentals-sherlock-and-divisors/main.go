package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Since sqrt(d) * sqrt(d) = d it's enough to find the divisors up
// to sqrt(d) and use them to calculate the subsequent divisors of d.
func divisors(d int) []int {
	var a []int = []int{}
	var i int = 0
	var s = int(math.Floor(math.Sqrt(float64(d))))

	// Find all divisors up to sqrt(d).
	for c := 1; c <= s; c++ {
		if d%c == 0 {
			a = append(a, c)
			i += 1
		}
	}

	// Omit the last divisor from the loop below when it equals to sqrt(d)
	// as the division would result in itself, adding a duplicate to the list.
	if a[i-1]*a[i-1] == d {
		i -= 1
	}

	// Calculate subsequent divisors by using those already found.
	for j := 0; j < i; j++ {
		a = append(a, d/a[i-j-1])
	}

	return a
}

func sherlockAndDivisors(n int) int {
	var a []int = divisors(n)
	var c int = 0
	var l int = len(a)

	for i := 0; i < l; i++ {
		if a[i]%2 == 0 {
			c += 1
		}
	}

	return c
}

func main() {
	in := bufio.NewReader(os.Stdin)
	t := parseInt64(readLine(in))
	for ; t > 0; t-- {
		n := parseInt(readLine(in))
		fmt.Printf("%d\n", sherlockAndDivisors(n))
	}
}

func parseInt64Seq(s string) []int64 {
	v := strings.Split(strings.TrimSpace(s), " ")
	r := []int64{}
	for i := range v {
		r = append(r, parseInt64(v[i]))
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
