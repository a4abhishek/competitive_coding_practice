package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrQueueUnderflow = errors.New("queue underflow")
)

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Queue []*TreeNode

func (q *Queue) Enqueue(n *TreeNode) {
	if q == nil {
		*q = []*TreeNode{}
	}

	*q = append(*q, n)
}

func (q *Queue) Dequeue() (*TreeNode, error) {
	if len(*q) == 0 {
		return nil, ErrQueueUnderflow
	}

	n := []*TreeNode(*q)[0]
	*q = []*TreeNode(*q)[1:]

	return n, nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) BFS() []*int {
	l := []*int{}

	if n == nil {
		return l
	}

	q := Queue{}
	q.Enqueue(n)

	for {
		x, err := q.Dequeue()
		if err != nil {
			break
		}

		if x == nil {
			l = append(l, nil)
			continue
		}

		l = append(l, &x.Val)

		if x.Left != nil || x.Right != nil {
			q.Enqueue(x.Left)
			q.Enqueue(x.Right)
		}
	}

	edge := len(l)
	for edge > 0 && l[edge-1] == nil {
		edge--
	}

	return l[:edge]
}

func (n *TreeNode) Deapth() int {
	if n == nil {
		return 0
	}

	lDeapth := n.Left.Deapth()
	rDeapth := n.Right.Deapth()

	return max(lDeapth, rDeapth) + 1
}

func sortedArrayToBST(l []int) *TreeNode {
	if len(l) == 0 {
		return nil
	}

	middleIndex := len(l) / 2

	tree := &TreeNode{}
	tree.Val = l[middleIndex]
	tree.Left = sortedArrayToBST(l[:middleIndex])
	tree.Right = sortedArrayToBST(l[middleIndex+1:])

	return tree
}

func main() {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	bfs := tree.BFS()

	fmt.Println(len(bfs))
	fmt.Println(bfs)

	jsonBytes, err := json.Marshal(bfs)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
