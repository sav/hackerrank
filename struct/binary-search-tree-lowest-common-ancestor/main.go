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
	return fmt.Sprintf("(%v%s%s)", t.Data, left, right)
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

type LCASearch[T any] struct {
	v1      T
	v2      T
	v1Found bool
	v2Found bool
}

func (s LCASearch[T]) found() bool {
	return s.v1Found && s.v2Found
}

func (s LCASearch[T]) String() string {
	return fmt.Sprintf("<%v/%v, %v/%v>", s.v1, s.v1Found, s.v2, s.v2Found)
}

func (t *BinaryTree) LCA(v1 int64, v2 int64) (*BinaryTree, LCASearch[int64]) {
	s := LCASearch[int64]{v1, v2, false, false}

	if t.Data == v1 {
		s.v1Found = true
	}

	if t.Data == s.v2 {
		s.v2Found = true
	}

	if !s.found() {
		if t.Left != nil {
			left, s1 := t.Left.LCA(v1, v2)
			if s1.found() {
				return left, s1
			} else if s1.v1Found {
				s.v1Found = true
			} else if s1.v2Found {
				s.v2Found = true
			}
		}
	}

	if !s.found() {
		if t.Right != nil {
			right, s2 := t.Right.LCA(v1, v2)
			if s2.found() {
				return right, s2
			} else if s2.v1Found {
				s.v1Found = true
			} else if s2.v2Found {
				s.v2Found = true
			}
		}
	}

	return t, s
}

func LCA(t *BinaryTree, v int64, w int64) *BinaryTree {
	a, s := t.LCA(v, w)
	if !s.found() {
		return nil
	}
	return a
}

func main() {
	in := bufio.NewReader(os.Stdin)
	parseInt64(readLine(in))
	v := parseInt64Seq(readLine(in))
	t := NewBinaryTree(v)
	a := parseInt64Seq(readLine(in))
	fmt.Printf("%v\n", LCA(t, a[0], a[1]).Data) // return value is never nil
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
