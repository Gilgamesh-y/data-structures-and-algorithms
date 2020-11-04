package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	RootNode TreeNode
}

type TreeNode struct {
	Value int
	LeftNode *TreeNode
	RightNode *TreeNode
}

var tree *Tree

/**
              5
          3       7
      1      4 6      9
         2        8       10
 */

func main() {
	arr := []int{5,3,1,7,9,2,4,6,8,10}
	for _, val := range arr {
		CreateTree(val)
	}
	println("先序遍历begin-------")
	preTraverseBTree(tree.RootNode)
	println("先序遍历end-------")

	println("中序遍历begin-------")
	inTraverseBTree(tree.RootNode)
	println("中序遍历end-------")

	println("后序遍历begin-------")
	nextTraverseBTree(tree.RootNode)
	println("后序遍历end-------")

	println("树的深度-------"+strconv.Itoa(getHeight(tree.RootNode)))
}

func CreateTree(value int) {
	if tree == nil {
		tree = &Tree{TreeNode{Value: value}}
	} else {
		currentNode := &tree.RootNode
		for currentNode != nil {
			if value > currentNode.Value {
				if currentNode.RightNode == nil {
					currentNode.RightNode = &TreeNode{Value: value}
					break
				} else {
					currentNode = currentNode.RightNode
				}
			}
			if value <= currentNode.Value {
				if currentNode.LeftNode == nil {
					currentNode.LeftNode = &TreeNode{Value: value}
					break
				} else {
					currentNode = currentNode.LeftNode
				}
			}
		}
	}
}

/**
 * 先序遍历
 */
func preTraverseBTree(node TreeNode) {
	if node != (TreeNode{}) {
		fmt.Println(node.Value)
		if node.LeftNode != nil {
			preTraverseBTree(*node.LeftNode)
		}
		if node.RightNode != nil {
			preTraverseBTree(*node.RightNode)
		}
	}
}

/**
 * 中序遍历
 */
func inTraverseBTree(node TreeNode) {
	if node != (TreeNode{}) {
		if node.LeftNode != nil {
			inTraverseBTree(*node.LeftNode)
		}
		fmt.Println(node.Value)
		if node.RightNode != nil {
			inTraverseBTree(*node.RightNode)
		}
	}
}

/**
 * 后序遍历
 */
func nextTraverseBTree(node TreeNode) {
	if node != (TreeNode{}) {
		if node.LeftNode != nil {
			nextTraverseBTree(*node.LeftNode)
		}
		if node.RightNode != nil {
			nextTraverseBTree(*node.RightNode)
		}
		fmt.Println(node.Value)
	}
}

func getHeight(node TreeNode) int {
	if node == (TreeNode{}) {
		return 0
	}
	leftHeight := 0
	rightHeight := 0
	if node.LeftNode != nil {
		leftHeight = getHeight(*node.LeftNode)
	}
	if node.RightNode != nil {
		rightHeight = getHeight(*node.RightNode)
	}
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1
}