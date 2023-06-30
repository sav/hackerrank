package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const Divisor = 1234567

func connectingTowns(roads []string) uint64 {
	var ret uint64 = 1
	for _, s := range roads {
		n, err := strconv.ParseInt(s, 10, 64)
		checkError(err)
		ret = (ret * uint64(n)) % Divisor
	}
	return ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tmp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tmp)

	for it := 0; it < t; it++ {
		_, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		roads := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		fmt.Printf("%d\n", connectingTowns(roads))
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
