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

/*
 *  To solve this problem, we need to determine:
 *
 *  1. How much the overlapping square decreases given a time displacement.
 *  2. An expression in terms of time `t' that has a real solution.
 *
 *  Step 1 will be determined by the Pythagorean Theorem, where the
 *  distance traveled is the hypotenuse of the triangle, while at the
 *  same time, it has equal-sized legs (cathetus) because they are the
 *  sides of the square. Therefore, they are of the same
 *  length. Furthermore, the distance traveled can be described in
 *  terms of velocity and time as `h = s * t', resulting in:
 *
 *     c = sqrt((s * t)²/2)
 *
 *  To obtain the expression from Step 2, we combine the above
 *  expression with the equation for the area of the square
 *  `q = (l - c)²' and isolate the unknown variable `t':
 *
 *    q = (l - sqrt((s * t)²/2))²
 *    t = sqrt(2 * (l - sqrt(q))²)/s
 */

func sherlockAndMovingTiles(l float64, s float64, q float64) float64 {
	return math.Sqrt(2*math.Pow(l-math.Sqrt(q), 2)) / s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	parms := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	l := parseFloat64(parms[0])

	s1 := parseFloat64(parms[1])

	s2 := parseFloat64(parms[2])

	s := float64(math.Abs(float64(s1 - s2)))

	t := int(parseFloat64(readLine(reader)))

	for it := 0; it < t; it++ {
		q := parseFloat64(readLine(reader))
		fmt.Printf("%.4f\n", sherlockAndMovingTiles(l, s, q))
	}
}

func parseFloat64(s string) float64 {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	checkError(err)
	return float64(n)
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
