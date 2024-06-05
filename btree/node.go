package main

// The order of a B tree is the minimum number of records a node can sustain. Except
// for the root, who can have a minimum of one record. Every node of a tree can sustain
// at least the order divided by two and rounded down.
// i.e.: The nodes of an order 5 tree can sustain at least 2 records (5/2 -> 2.5 -> 2).
// The number of childs a node can have is always the number of records of the node plus one.
// i.e.: If the node has 4 records, it will necessarily point out to 4 childrens.
// Order and leaf

const _order = 5

type Node struct {
	key   []int
	child []Node
}

func NewNode() Node {
	return Node{
		key:   make([]int, 0, _order-1),
		child: make([]Node, 0, _order),
	}
}

func (n Node) SearchRecursive(target int) Node {
	next := len(n.child) - 1
	nextSelected := false
	for i, k := range n.key {
		if k == target {
			return n
		}
		if !nextSelected && target < k && len(n.child) > 0 {
			next = i
			nextSelected = true
		}
	}

	if next < 0 {
		return Node{}
	}

	return n.child[next].SearchRecursive(target)
}

func (n Node) SearchInPlace(target int) Node {
	searchNode := n
	for {
		next := len(searchNode.child) - 1
		nextSelected := false
		for i, k := range searchNode.key {
			if k == target {
				return searchNode
			}

			if !nextSelected && target < k && len(searchNode.child) > 0 {
				next = i
				nextSelected = true
			}
		}

		if next < 0 {
			return Node{}
		}

		searchNode = searchNode.child[next]
	}
}

func (n *Node) Insert(key int) error {

	return nil
}
