package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Tree struct {
	RootNode *TreeNode
}

type TreeNode struct {
	Value int
	LeftNode *TreeNode
	RightNode *TreeNode
}

/**
              5
          3       7
      1      4 6      9
         2        8       10
 */

func main() {
	tree := &Tree{}
	arr := []int{5,3,1,7,9,2,4,6,8,10}
	for _, val := range arr {
		tree.CreateTree(val)
	}
	println("先序遍历begin-------")
	tree.RootNode.PreTraverseBTree()
	println("先序遍历end-------")

	println("中序遍历begin-------")
	tree.RootNode.InTraverseBTree()
	println("中序遍历end-------")

	println("后序遍历begin-------")
	tree.RootNode.NextTraverseBTree()
	println("后序遍历end-------")

	println("树的深度-------"+strconv.Itoa(tree.getHeight(tree.RootNode)))

	println("查找begin------")
	pnode, node, _ := tree.Search(5, tree.RootNode, tree.RootNode, false)
	pnodestr, _ := json.Marshal(pnode)
	println("父节点"+string(pnodestr))
	nodestr, _ := json.Marshal(node)
	println("子节点"+string(nodestr))
	println("查找end------")

	println("删除begin-------")
	tree.DelNode(3)
	println("删除end-------")
	nodestr, _ = json.Marshal(tree.RootNode)
	println("删除后的树"+string(nodestr))
}

func (tree *Tree) CreateTree(value int) {
	if tree.RootNode == nil {
		tree.RootNode =&TreeNode{Value: value}
	} else {
		currentNode := tree.RootNode
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
func (node *TreeNode) PreTraverseBTree() {
	if node != nil {
		fmt.Println(node.Value)
		if node.LeftNode != nil {
			node.LeftNode.PreTraverseBTree()
		}
		if node.RightNode != nil {
			node.RightNode.PreTraverseBTree()
		}
	}
}

/**
 * 中序遍历
 */
func (node *TreeNode) InTraverseBTree() {
	if node != nil {
		if node.LeftNode != nil {
			node.LeftNode.InTraverseBTree()
		}
		fmt.Println(node.Value)
		if node.RightNode != nil {
			node.RightNode.InTraverseBTree()
		}
	}
}

/**
 * 后序遍历
 */
func (node *TreeNode) NextTraverseBTree() {
	if node != nil {
		if node.LeftNode != nil {
			node.LeftNode.NextTraverseBTree()
		}
		if node.RightNode != nil {
			node.RightNode.NextTraverseBTree()
		}
		fmt.Println(node.Value)
	}
}

/**
 * 获取树的高度
 */
func (tree *Tree) getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := 0
	rightHeight := 0
	if node.LeftNode != nil {
		leftHeight = tree.getHeight(node.LeftNode)
	}
	if node.RightNode != nil {
		rightHeight = tree.getHeight(node.RightNode)
	}
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1
}

func (tree *Tree) Search(value int, parrentNode *TreeNode, node *TreeNode, isLeft bool) (*TreeNode, *TreeNode, bool) {
	if node == nil {
		return parrentNode, nil, isLeft
	}
	if parrentNode == nil {
		return nil, nil, isLeft
	}
	if value < node.Value {
		isLeft = true
		return tree.Search(value, node, node.LeftNode, isLeft)
	}
	if value > node.Value {
		isLeft = false
		return tree.Search(value, node, node.RightNode, isLeft)
	}

	return parrentNode, node, isLeft
}

func (node *TreeNode) HasOneChild() (*TreeNode, bool) {
	isLeft := false
	if node.LeftNode != nil && node.RightNode == nil {
		isLeft = true
		return node.LeftNode, isLeft
	}
	if node.RightNode != nil && node.LeftNode == nil {
		return node.RightNode, isLeft
	}
	return nil, isLeft
}

func (node *TreeNode) HasTwoChild() bool {
	if node.LeftNode != nil && node.RightNode != nil {
		return true
	}
	return false
}

func (node *TreeNode) GetLeftMaxNode(parrentNode *TreeNode,depth int) (*TreeNode, *TreeNode) {
	if node.LeftNode == nil && depth == 0 {
		return nil, node
	}
	if node.RightNode != nil {
		return node.RightNode.GetLeftMaxNode(node, depth + 1)
	}

	return parrentNode, node
}

func (tree *Tree) DelNode(value int) bool {
	pnode, node, isLeft := tree.Search(value, tree.RootNode, tree.RootNode, false)
	if node == nil {
		return false
	}
	// 有左子节点和右子节点
	if node.HasTwoChild() {
		lParrentMaxNode, lMaxNode := node.GetLeftMaxNode(node, 0)
		node.Value = lMaxNode.Value
		lParrentMaxNode.RightNode = nil
		return true
	}
	childNode, _ := node.HasOneChild()
	// 如果有子节点
	if childNode != nil {
		childNode = nil
		return false
	}
	// 如果没有子节点
	if childNode == nil {
		if isLeft {
			pnode.LeftNode = nil
			return true
		}
		pnode.RightNode = nil
		return true
	}
	// 只有左子节点或右子节点
	*node = *childNode
	return true
}