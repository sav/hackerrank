package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sumDigits(n int64) int64 {
	var s int64 = 0
	var d int64 = 10
	for {
		s += ((n % d) / (d / 10))
		if n/d == 0 {
			break
		}
		d *= 10
	}
	return s
}

func findBest(d []int64, s []int64) int64 {

	var bestIdx int64 = 0
	var bestSum int64 = 1

	for i := range d {
		if s[i] > bestSum {
			bestIdx = int64(i)
			bestSum = s[i]
		}
	}

	return d[bestIdx]
}

func bestDivisor(n int64) int64 {
	var d []int64 = []int64{}
	var s []int64 = []int64{}

	var len int64 = n
	var i int64 = 1

	for ; i <= len; i++ {
		if n%i == 0 {
			d = append(d, i)
			s = append(s, sumDigits(i))
		}
	}

	return findBest(d, s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := parseInt64(readLine(reader))
	fmt.Printf("%d\n", bestDivisor(n))
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
