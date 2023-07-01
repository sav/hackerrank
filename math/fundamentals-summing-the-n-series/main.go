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

var (
	Divisor = big.NewInt(1000000007)
)

/*
 * First step it to simplifly Tn = n^2 - (n-1)^2:
 * n^2 - (n - 1)^2 = n^2 - (n^2 - 2*n + 1) = 2n - 1
 *
 * The equation Sn = T1 + ... + Tn becomes:
 * for i := 1; i <= n; i++ { sum = (sum + 2*n - 1) % Divisor }
 *
 * By induction we can show that:
 * (2*n - 1) + (2*(n-1) - 1) + ... + 1 = 2*n^2 - n^2 = n^2
 *
 * We'd then have: n^2 % Divisor, but this is the same as:
 * ((n % Divisor) * n) % Divisor
 *
 * The latter being better since it fits in less bits.
 *
 * Unfortunately `(10^16 mod Divisor) * 10^16' does not fit
 * in 64-bit, so we need to use bignum.
 */

func summingTheNSeries(n int64) *big.Int {
	var r *big.Int = big.NewInt(n)
	r.Mod(r, Divisor)
	r.Mul(r, big.NewInt(n))
	r.Mod(r, Divisor)
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tmp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tmp)

	for it := 0; it < t; it++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		fmt.Printf("%s\n", summingTheNSeries(int64(n)))
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
