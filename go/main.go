package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Node struct {
	X     int
	Y     int
	Left  *Node
	Right *Node
}

func NewNode(v int) *Node {
	return &Node{
		X: v,
		Y: rand.Int(),
	}
}

func (n *Node) split(value int) (lower, equal, greater *Node) {
	var equalGreater *Node
	lower, equalGreater = n.splitBinary(value)
	equal, greater = equalGreater.splitBinary(value + 1)
	return
}

func (n *Node) splitBinary(value int) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}

	if n.X < value {
		splitPair0, splitPair1 := n.Right.splitBinary(value)
		n.Right = splitPair0
		return n, splitPair1
	}

	splitPair0, splitPair1 := n.Left.splitBinary(value)
	n.Left = splitPair1
	return splitPair0, n
}

type Tree struct {
	Root *Node
}

func (t *Tree) HasValue(v int) bool {
	lower, equal, greater := t.Root.split(v)
	res := equal != nil
	t.Root = merge3(lower, equal, greater)
	return res
}

func (t *Tree) Insert(v int) error {
	lower, equal, greater := t.Root.split(v)
	if equal == nil {
		equal = NewNode(v)
	}
	t.Root = merge3(lower, equal, greater)
	return nil
}

func (t *Tree) Erase(v int) error {
	lower, _, greater := t.Root.split(v)
	t.Root = merge(lower, greater)
	return nil
}

func merge(lower, greater *Node) *Node {
	if lower == nil {
		return greater
	}

	if greater == nil {
		return lower
	}

	if lower.Y < greater.Y {
		right := merge(lower.Right, greater)
		lower.Right = right
		return lower
	}
	left := merge(lower, greater.Left)
	greater.Left = left
	return greater
}

func merge3(lower, equal, greater *Node) *Node {
	return merge(merge(lower, equal), greater)
}

func main() {
	t := &Tree{}

	cur := 5
	res := 0

	for i := 1; i < 1000000; i++ {
		a := i % 3
		cur = (cur*57 + 43) % 10007
		if a == 0 {
			t.Insert(cur)
		} else if a == 1 {
			t.Erase(cur)
		} else if a == 2 {
			has := t.HasValue(cur)
			if has {
				res += 1
			}
		}
	}

	fmt.Println(res)
}
