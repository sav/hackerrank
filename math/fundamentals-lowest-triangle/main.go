package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/* a = h * b / 2
 * h = 2 * a / b
 */

func lowestTriangle(b uint64, a uint64) uint64 {
	h := 2 * a / b
	if h*b/2 < a {
		h += 1
	}
	return h
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	v := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	b, err := strconv.ParseInt(v[0], 10, 64)
	checkError(err)

	a, err := strconv.ParseInt(v[1], 10, 64)
	checkError(err)

	fmt.Printf("%d\n", lowestTriangle(uint64(b), uint64(a)))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
