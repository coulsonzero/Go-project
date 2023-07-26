package main

import "math"

/**
 * 给定一个 N 叉树，找到其最大深度。
 * 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
 * 输入：root = [1,null,3,2,4,null,5,6]
 * 输出：3
 */

/**
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil {
		return 1
	}

	depth := 0
	for _, c := range root.Children {
		depth = int(math.Max(float64(depth), float64(maxDepth(c))))
	}
	return depth + 1
}
