package main

import (
	"fmt"
)

type Node struct {
	UsedKeysNum int8
	Keys        [3]int
	Nodes       [4]*Node
}

func NewNode() *Node {
	return &Node{
		UsedKeysNum: 0,
		Keys:        [3]int{0, 0, 0},
		Nodes:       [4]*Node{nil, nil, nil, nil},
	}
}

func (n *Node) Shift2Right() {
	for i := 2; i >= 1; i-- {
		n.Keys[i] = n.Keys[i-1]
		n.Nodes[i] = n.Nodes[i-1]
	}
	n.Nodes[0] = nil
}

func (n *Node) Add(newVal int) {
	if n.UsedKeysNum == 0 {
		n.Keys[0] = newVal
		n.UsedKeysNum++
		return
	}
	if n.UsedKeysNum == 1 {
		if n.Keys[0] < newVal {
			n.Keys[1] = newVal
		} else {
			n.Shift2Right()
			n.Keys[0] = newVal
		}
		n.UsedKeysNum++
		return
	}
	if n.UsedKeysNum == 2 {
		if n.Keys[1] < newVal {
			n.Keys[2] = newVal
		} else if n.Keys[0] < newVal {
			n.Keys[2] = n.Keys[1]
			n.Nodes[2] = n.Nodes[1]
			n.Keys[1] = newVal
			n.Nodes[1] = nil
		} else {
			n.Shift2Right()
			n.Keys[0] = newVal
		}
		n.UsedKeysNum++
		return
	}
	for i := 0; i < 3; i++ {
		if newVal < n.Keys[i] {
			if n.Nodes[i] == nil {
				n.Nodes[i] = NewNode()
			}
			n.Nodes[i].Add(newVal)
			return
		}
	}
	if n.Nodes[3] == nil {
		n.Nodes[3] = NewNode()
	}
	n.Nodes[3].Add(newVal)
}

func main() {
	n := NewNode()
	n.Add(3)
	fmt.Println(n.Keys, n.UsedKeysNum)
	n.Add(1)
	fmt.Println(n.Keys, n.UsedKeysNum)
	n.Add(2)
	fmt.Println(n.Keys, n.UsedKeysNum)
}
