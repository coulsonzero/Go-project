package main

import "fmt"

// 二叉树的后序遍历
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

func levelOrder(root *TreeNode) (res [][]int) {
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}

	}

	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func main() {
	// [3,9,20,null,null,15,7]
	node5 := TreeNode{7, nil, nil}
	node4 := TreeNode{15, nil, nil}
	node3 := TreeNode{20, &node4, &node5}
	node2 := TreeNode{9, nil, nil}
	node1 := TreeNode{3, &node2, &node3}

	fmt.Printf("res is: %v \n", levelOrder(&node1))
	//  [[3] [9 20] [15 7]]
}
