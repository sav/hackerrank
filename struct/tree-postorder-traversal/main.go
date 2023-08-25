package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//// Queue

type QueueError int8

const (
	QueueErrorNil QueueError = iota
	QueueErrorEmpty
)

var queueErrorString = map[QueueError]string{
	QueueErrorNil:   "Null pointer",
	QueueErrorEmpty: "Queue is empty",
}

func (err QueueError) String() string {
	return queueErrorString[err]
}

func (err QueueError) Error() string {
	return "Queue error: " + err.String()
}

type Queue[T any] struct {
	list []T
}

func (q Queue[T]) String() string {
	var s = "["
	for _, e := range q.list {
		s += fmt.Sprintf("%v ", e)
	}
	if len(q.list) > 0 {
		s = s[:len(s)-1]
	}
	s += "]"
	return s
}

func (q *Queue[T]) Size() int {
	return len(q.list)
}

func (q *Queue[T]) Front() *T {
	if q == nil {
		return nil
	}
	if q.Size() == 0 {
		return nil
	}
	return &q.list[0]
}

func (q *Queue[T]) Back() *T {
	if q == nil {
		return nil
	}
	if q.Size() == 0 {
		return nil
	}
	return &q.list[len(q.list)-1]
}

func (q *Queue[T]) Push(e T) error {
	if q == nil {
		return QueueErrorNil
	}
	q.list = append(q.list, e)
	return nil
}

func (q *Queue[T]) Pop() error {
	if q == nil {
		return QueueErrorNil
	}
	if len(q.list) == 0 {
		return QueueErrorEmpty
	}
	q.list = q.list[1:]
	return nil
}

//// BinaryTree

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

func (t *BinaryTree) Height() int {
	if t == nil || (t.Left == nil && t.Right == nil) {
		return 0
	}
	lHeight, rHeight := t.Left.Height(), t.Right.Height()
	if lHeight > rHeight {
		return 1 + lHeight
	}
	return 1 + rHeight
}

func NewBinaryTree(elems []int64) *BinaryTree {
	var tree = &BinaryTree{}
	for _, elem := range elems {
		tree.Insert(elem)
	}
	return tree
}

//// BFS

type bfSearch struct {
	node  *BinaryTree
	level int
}

func (t *BinaryTree) doBFS(queue *Queue[bfSearch], callback func(*BinaryTree, int) bool) {
	if queue.Size() == 0 {
		return
	}
	var search = queue.Front()
	if !callback(search.node, search.level) {
		return
	}
	if search.node.Left != nil {
		queue.Push(bfSearch{search.node.Left, search.level - 1})
	}
	if search.node.Right != nil {
		queue.Push(bfSearch{search.node.Right, search.level + 1})
	}
	queue.Pop()
	t.doBFS(queue, callback)
}

// BFS visits every node at each level of the tree before moving to the next level.
// It calls the callback function for each visited node. If the callback returns false, the function
// stops iterating and returns. Otherwise, it continues iterating until reaching the end.
func (t *BinaryTree) BFS(callback func(*BinaryTree, int) bool) {
	var queue = &Queue[bfSearch]{}
	queue.Push(bfSearch{t, 0})
	t.doBFS(queue, callback)
}

func (t *BinaryTree) InOrder(callback func(*BinaryTree) bool) {
	if t.Left != nil {
		t.Left.InOrder(callback)
	}
	callback(t)
	if t.Right != nil {
		t.Right.InOrder(callback)
	}
}

func (t *BinaryTree) PostOrder(callback func(*BinaryTree) bool) {
	if t.Left != nil {
		t.Left.PostOrder(callback)
	}
	if t.Right != nil {
		t.Right.PostOrder(callback)
	}
	callback(t)
}

//// Problem Solution

func (t *BinaryTree) PostOrderTraversal() {
	t.PostOrder(func(node *BinaryTree) bool {
		fmt.Printf("%d ", node.Data)
		return true
	})
}

//// main

func main() {
	in := bufio.NewReader(os.Stdin)
	parseInt64(readLine(in))
	v := parseInt64Seq(readLine(in))
	t := NewBinaryTree(v)
	t.PostOrderTraversal()
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
