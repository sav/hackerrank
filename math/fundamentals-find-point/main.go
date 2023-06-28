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
 * Complete the 'findPoint' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER px
 *  2. INTEGER py
 *  3. INTEGER qx
 *  4. INTEGER qy
 */

func findPoint(px int32, py int32, qx int32, qy int32) {
	fmt.Printf("%d %d\n", 2*qx-px, 2*qy-py)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for nItr := 0; nItr < int(n); nItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		pxTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		px := int32(pxTemp)

		pyTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		py := int32(pyTemp)

		qxTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		qx := int32(qxTemp)

		qyTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		qy := int32(qyTemp)

		findPoint(px, py, qx, qy)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
