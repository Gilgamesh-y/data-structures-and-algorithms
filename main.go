package main

type Tree struct {
	TreeNode
}

type TreeNode struct {
	Value int
	LeftNode *TreeNode
	RightNode *TreeNode
}

var tree Tree

func main() {
	arr := []int{1,3,5,7,9,2,4,6,8,10}
}

func (t *TreeNode) SetLeftNode(node *TreeNode) {
	t.LeftNode = node
}

func (t *TreeNode) SetRightNode(node *TreeNode) {
	t.RightNode = node
}

func (t *Tree) GetTreeNode() TreeNode {
	return t.TreeNode
}

func (t *Tree) SetTreeNode(node TreeNode) {
	t.TreeNode = node
}


func CreateTree(tree *Tree, value int) {
	if tree.GetTreeNode() == (TreeNode{}) {
		node := TreeNode{Value: value}
		tree.SetTreeNode(node)
	} else {
		node := tree.GetTreeNode()
		for node != (TreeNode{}) {
			
		}
	}
}