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

/*
 * Complete the 'handshake' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

var (
	Zero *big.Int = big.NewInt(0)
	One  *big.Int = big.NewInt(1)
	Two  *big.Int = big.NewInt(2)
)

// The number of combinations "n choose k" without repetition is:
//   C(n, k) = n! / (k! * (n - k)!)

func handshake(n int64) {
	if n < 2 {
		fmt.Println(0)
		return
	}

	var x *big.Int = big.NewInt(1)
	var i int64

	for i = 2; i <= n; i++ {
		x.Mul(x, big.NewInt(i))
	}

	var y *big.Int = new(big.Int).Div(x, big.NewInt(n))
	y.Div(y, big.NewInt(n-1)).Mul(y, Two)

	fmt.Printf("%s\n", x.Div(x, y))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tmp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tmp)

	for it := 0; it < t; it++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		handshake(int64(n))
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
