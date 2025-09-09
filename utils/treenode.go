package utils

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBinaryTree builds a binary tree from level-order representation where nils are represented by a special flag value.
// Use flag value math.MinInt to represent nils when building.
func BuildBinaryTree(values []int, nilFlag int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(values))
	for i, v := range values {
		if v == nilFlag {
			nodes[i] = nil
			continue
		}
		nodes[i] = &TreeNode{Val: v}
	}
	for i := 0; i < len(values); i++ {
		if nodes[i] == nil {
			continue
		}
		li := 2*i + 1
		ri := 2*i + 2
		if li < len(values) {
			nodes[i].Left = nodes[li]
		}
		if ri < len(values) {
			nodes[i].Right = nodes[ri]
		}
	}
	return nodes[0]
}
