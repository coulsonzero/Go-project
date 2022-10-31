package main

import (
	"fmt"
)

// Definition for a binary tree node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的前序遍历
// Input: root = [1,null,2,3]
// Output: [1,2,3]

// 递归
func preorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		order(node.Left)
		order(node.Right)
	}
	if root == nil {
		return nil
	}
	order(root)
	return
}

func main() {
	node3 := TreeNode{3, nil, nil}
	node2 := TreeNode{2, &node3, nil}
	root := TreeNode{1, nil, &node2}

	cur := preorderTraversal(&root)
	fmt.Println(cur) // 1 2 3
}
