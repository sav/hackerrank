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

func gameWithCells(n uint64, m uint64) uint64 {
	return uint64(math.Ceil(float64(n)/2.0) * math.Ceil(float64(m)/2.0))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	v := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n, err := strconv.ParseInt(v[0], 10, 64)
	checkError(err)

	m, err := strconv.ParseInt(v[1], 10, 64)
	checkError(err)

	fmt.Printf("%d\n", gameWithCells(uint64(n), uint64(m)))
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
