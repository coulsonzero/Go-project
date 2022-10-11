package main

import "fmt"

/**
 * Question: JZ55 二叉树的深度
 *
 * description: 输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度，根节点的深度视为 1 。
 * input : {1,2,3,4,5,nil,6,nil,nil,7}
 * output: 4
 */

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}

	return 1 + max(TreeDepth(pRoot.Left), TreeDepth(pRoot.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node6 := TreeNode{Val: 6, Left: nil, Right: nil}
	node5 := TreeNode{Val: 5, Left: nil, Right: &node7}
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := TreeNode{Val: 3, Left: nil, Right: &node6}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	root := TreeNode{Val: 1, Left: &node2, Right: &node3}
	fmt.Println(TreeDepth(&root)) // 4
}
