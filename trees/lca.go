package trees

import "github.com/Rehanqr7/dsa-go/utils"

/*
Lowest Common Ancestor of a Binary Tree (Medium)

Time Complexity: O(n)
Space Complexity: O(h)
*/

func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
