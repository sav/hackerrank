package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func diwaliLight(i int64) int64 {
	n := big.NewInt(i)
	return n.Exp(big.NewInt(2), n, big.NewInt(100000)).Int64() - 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	t := parseInt64(readLine(in))
	for ; t > 0; t-- {
		n := parseInt64(readLine(in))
		fmt.Printf("%d\n", diwaliLight(n))
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
