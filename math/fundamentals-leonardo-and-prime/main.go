package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// We only need to multiply the first 16 primes to exceed
// the constraint of 10^18.

var Primes []uint64 = []uint64{
	2,
	3,
	5,
	7,
	11,
	13,
	17,
	19,
	23,
	29,
	31,
	37,
	41,
	43,
	47,
	53,
}

func leonardoAndPrime(n uint64) int {
	var mul uint64 = 1
	var cnt = 0
	for {
		mul *= Primes[cnt]
		if mul > n {
			break
		}
		cnt += 1
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tmp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tmp)

	for it := 0; it < t; it++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		fmt.Printf("%d\n", leonardoAndPrime(uint64(n)))
	}
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
