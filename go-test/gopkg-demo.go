package main

/*
    2
   /  \
  3   9
 /   /
4   7
*/

var res [][]int

/*
[
	[2],
	[3, 9],


]

*/

func order(root *TreeNode) {
	if root == nil {
		return
	}
	var temp []int
	var level int
	res[level] = append(res[level], root.Val)

	if root.Left != nil {
		temp = append(temp, root.Left.Val)
	}
	if root.Right != nil {
		temp = append(temp, root.Right.Val)
	}
	res = append(res, temp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node5 := TreeNode{7, nil, nil}
	node4 := TreeNode{4, nil, nil}
	node3 := TreeNode{9, &node5, nil}
	node2 := TreeNode{3, &node4, nil}
	root := TreeNode{2, &node2, &node3}

	order(&root)
}


