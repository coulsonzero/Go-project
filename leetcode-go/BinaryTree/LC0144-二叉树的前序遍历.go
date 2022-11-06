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

/*
先序遍历是中左右，中序遍历是左中右， 后序遍历是左右中
*/

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
	return res
}

// 迭代
func preorderTraversal2(root *TreeNode) (vals []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func main() {
	node3 := TreeNode{3, nil, nil}
	node2 := TreeNode{2, &node3, nil}
	root := TreeNode{1, nil, &node2}

	cur := preorderTraversal(&root)
	fmt.Println(cur) // 1 2 3
}
