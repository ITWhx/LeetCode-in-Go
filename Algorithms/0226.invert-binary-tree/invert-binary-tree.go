package problem0226

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//不能这样写，这样翻转子节点时，会改变root结构
	//root.Left = invertTree(root.Right)
	//root.Right = invertTree(root.Left)

	r := invertTree(root.Right)
	l := invertTree(root.Left)

	root.Left = r
	root.Right = l
	return root
}
