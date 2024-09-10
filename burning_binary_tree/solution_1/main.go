package main

import (
	"errors"
	"fmt"
)

// Queue DataStructure
type Queue[T any] struct {
	zeroValue T
	s         []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		s: []T{},
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.s = append(q.s, n)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.s) == 0 {
		return q.zeroValue, errors.New("Queue Underflow")
	}

	n := q.s[0]
	q.s = q.s[1:]

	return n, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.s) == 0 {
		return q.zeroValue, errors.New("Queue Underflow")
	}

	return q.s[0], nil
}

func (q *Queue[T]) Len() int {
	return len(q.s)
}

// Tree DataStructure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func getParentMap(currentNode, parentNode *TreeNode, parentMap map[*TreeNode]*TreeNode, target int) *TreeNode {
	if currentNode == nil {
		return nil
	}

	parentMap[currentNode] = parentNode

	var targetNode *TreeNode
	if currentNode.Left != nil {
		node := getParentMap(currentNode.Left, currentNode, parentMap, target)
		if node != nil {
			targetNode = node
		}
	}

	if currentNode.Right != nil {
		node := getParentMap(currentNode.Right, currentNode, parentMap, target)
		if node != nil {
			targetNode = node
		}
	}

	if currentNode.Val == target {
		return currentNode
	}

	return targetNode
}

func burn(targetNode *TreeNode, parentMap map[*TreeNode]*TreeNode) int {
	if targetNode == nil {
		return 0
	}

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(targetNode)

	// fmt.Printf("Initial Queue Length: %d\n", queue.Len())

	visited := map[*TreeNode]bool{}
	burnTime := -1 // As burning starts at 0th minute. i.e. first node will burn when burnTime is 0
	for qLen := queue.Len(); qLen != 0; qLen = queue.Len() {
		burnTime++ // Progressing the Burn

		// fmt.Printf("Current Queue Length: %d\n", queue.Len())
		// fmt.Printf("Time: %d. Burning nodes: ", burnTime)
		// for i := range qLen {
		// 	fmt.Printf("%d ", queue.s[i].Val)
		// }
		// fmt.Println()

		for range qLen {
			node, err := queue.Dequeue()
			if err != nil {
				break
			}
			if visited[node] {
				continue
			}
			visited[node] = true

			// fmt.Printf("Dequeued %d\n", node.Val)
			if node.Left != nil && !visited[node.Left] {
				// fmt.Printf("Enquiuing left node of %d: %d\n", node.Val, node.Left.Val)
				queue.Enqueue(node.Left)
			}
			if node.Right != nil && !visited[node.Right] {
				// fmt.Printf("Enquiuing right node of %d: %d\n", node.Val, node.Right.Val)
				queue.Enqueue(node.Right)
			}
			if parentMap[node] != nil && !visited[parentMap[node]] {
				// fmt.Printf("Enquiuing parent of %d: %d\n", node.Val, parentMap[node].Val)
				queue.Enqueue(parentMap[node])
			}
		}
	}

	return burnTime
}

func amountOfTime(root *TreeNode, start int) int {
	// Create Parent Map, and find target-node
	parentMap := map[*TreeNode]*TreeNode{}
	targetNode := getParentMap(root, nil, parentMap, start) // Parent of root is nil, and current parentMap is empty. It also returns the targetNode

	// for node, parent := range parentMap {
	// 	var p int
	// 	if parent != nil {
	// 		p = parent.Val
	// 	}
	// 	fmt.Printf("Parent of %d is %d\n", node.Val, p)
	// }

	// Proceed with Burn
	return burn(targetNode, parentMap)
}

func main() {
	tree := NewNode(1)
	tree.Left = NewNode(5)
	tree.Left.Right = NewNode(4)
	tree.Left.Right.Left = NewNode(9)
	tree.Left.Right.Right = NewNode(2)
	tree.Right = NewNode(3)
	tree.Right.Left = NewNode(10)
	tree.Right.Right = NewNode(6)

	fmt.Println(amountOfTime(tree, 3))
}
