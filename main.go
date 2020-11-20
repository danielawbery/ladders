package main

import "fmt"

// TreeNode is a data structure that has Val of type int and Left and Right brances that point to other TreeNodes
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(value int) *TreeNode {
	tn := TreeNode{value: value}
	return &tn
}

func mergeTrees(nodes1 []*TreeNode, nodes2 []*TreeNode) []*TreeNode {
	merged := make([]*TreeNode, 0)
	merged = append(merged, nodes1[0])

	return merged
}

func main() {
	// Instantiate Tree 1 input with their values and relationships
	tree1Node5 := newTreeNode(5)
	tree1Node3 := newTreeNode(3)
	tree1Node2 := newTreeNode(2)
	tree1Node1 := newTreeNode(1)

	tree1Node1.left = tree1Node3
	tree1Node1.right = tree1Node2
	tree1Node3.left = tree1Node5

	// Creating slice of TreeNode, with root first and assume ordering left to right
	treeNodes1 := []*TreeNode{tree1Node1, tree1Node3, tree1Node2, tree1Node5}
	fmt.Println(treeNodes1)

	// Instantiate Tree 2 input with their values and relationships
	tree2Node4 := newTreeNode(4)
	tree2Node7 := newTreeNode(7)
	tree2Node1 := newTreeNode(1)
	tree2Node3 := newTreeNode(3)
	tree2Node2 := newTreeNode(2)

	tree2Node2.left = tree2Node1
	tree2Node2.right = tree2Node3
	tree2Node1.right = tree2Node4
	tree2Node3.right = tree2Node7

	// Creating slice of TreeNode, with root first and assume ordering left to right
	treeNodes2 := []*TreeNode{tree2Node2, tree2Node1, tree2Node3, tree2Node4, tree2Node7}
	fmt.Println(treeNodes2)

	merged := mergeTrees(treeNodes1, treeNodes2)
	fmt.Println(merged)
}
