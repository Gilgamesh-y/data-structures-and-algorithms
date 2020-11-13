package main

import (
	"fmt"
)

type Tree struct {
	RootNode *TreeNode
}

type TreeNode struct {
	Value int
	LeftNode *TreeNode
	RightNode *TreeNode
}

func main() {
	arr := []int{5,3,1,7,9,2,4,6,8,10}

	tree := &Tree{}

	for _, val := range arr {
		tree.Insert(val)
	}

	fmt.Printf("%v\n", tree.Search(4))
}

func (tree *Tree) Search(value int) *TreeNode {
	return _search(value, tree.RootNode)
}

func _search(value int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Value {
		return _search(value, node.LeftNode)
	} else if value > node.Value {
		return _search(value, node.RightNode)
	} else {
		return node
	}
}

func (tree *Tree) Insert(value int) {
	if tree.RootNode == nil {
		tree.RootNode = &TreeNode{Value: value}
		return
	}
	_insert(value, tree.RootNode)
}

func _insert(value int, node *TreeNode) {
	if value < node.Value {
		if node.LeftNode != nil {
			_insert(value, node.LeftNode)
		} else {
			node.LeftNode = &TreeNode{Value:value}
		}
	} else {
		if node.RightNode != nil {
			_insert(value, node.RightNode)
		} else {
			node.RightNode = &TreeNode{Value:value}
		}
	}
}
