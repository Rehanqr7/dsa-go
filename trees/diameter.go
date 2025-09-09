package trees

import "github.com/Rehanqr7/dsa-go/utils"

/*
Diameter of Binary Tree (Easy/Medium)
Return the length of the diameter (number of edges in the longest path between any two nodes).

Time Complexity: O(n)
Space Complexity: O(h)
*/

func DiameterOfBinaryTree(root *utils.TreeNode) int {
	maxDiameter := 0
	var depth func(*utils.TreeNode) int
	depth = func(n *utils.TreeNode) int {
		if n == nil {
			return 0
		}
		l := depth(n.Left)
		r := depth(n.Right)
		if l+r > maxDiameter {
			maxDiameter = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	depth(root)
	return maxDiameter
}
