package Summary

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // 前序
		order(node.Left)
		order(node.Right)
	}
	if root == nil {
		return nil
	}
	order(root)
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		res = append(res, node.Val) // 中序
		order(node.Right)
	}
	if root == nil {
		return nil
	}
	order(root)
	return
}

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var order func(*TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		order(node.Right)
		res = append(res, node.Val) // 后序
	}
	if root == nil {
		return nil
	}
	order(root)
	return
}

// 层序遍历
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

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m1 := minDepth(root.Left)
	m2 := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	}
	return min(m1, m2) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && check(root1.Left, root2.Right) && check(root1.Right, root2.Left)
}
