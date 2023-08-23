package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BinaryTree struct {
	Data  int64
	Left  *BinaryTree
	Right *BinaryTree
}

func (t *BinaryTree) String() string {
	var left, right = "", ""
	if t.Left != nil {
		left = " " + t.Left.String()
	}
	if t.Right != nil {
		right = " " + t.Right.String()
	}
	return fmt.Sprintf("%v%s%s", t.Data, left, right)
}

func (t *BinaryTree) Insert(v int64) {
	if t.Data == 0 {
		t.Data = v
	} else if v <= t.Data {
		if t.Left == nil {
			t.Left = &BinaryTree{}
		}
		t.Left.Insert(v)
	} else {
		if t.Right == nil {
			t.Right = &BinaryTree{}
		}
		t.Right.Insert(v)
	}
}

func NewBinaryTree(elems []int64) *BinaryTree {
	var tree = &BinaryTree{}
	for _, elem := range elems {
		tree.Insert(elem)
	}
	return tree
}

func main() {
	in := bufio.NewReader(os.Stdin)
	parseInt64(readLine(in))
	v := parseInt64Seq(readLine(in))
	t := NewBinaryTree(v)
	fmt.Printf("%v\n", t)
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
