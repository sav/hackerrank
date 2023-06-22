package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b uint32
	fmt.Scanf("%v\n%v", &a, &b)
	fmt.Println(solveMeFirst(a, b))
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func scanInt(scanner *bufio.Scanner) (int, error) {
// 	ret := scanner.Scan()
// 	if ret {
// 		text := scanner.Text()
// 		return strconv.Atoi(text)
// 	}
// 	return -1, scanner.Err()
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	a, err := scanInt(scanner)
// 	if a < 1 || err != nil {
// 		fmt.Println("Error reading input:", err)
// 		os.Exit(1)
// 	}
// 	b, err := scanInt(scanner)
// 	if b < 1 || err != nil {
// 		fmt.Println("Error reading input:", err)
// 		os.Exit(2)
// 	}
// 	fmt.Println(a + b)
// }
