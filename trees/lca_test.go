package trees

import (
	"testing"

	"github.com/Rehanqr7/dsa-go/utils"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &utils.TreeNode{Val: 3}
	root.Left = &utils.TreeNode{Val: 5}
	root.Right = &utils.TreeNode{Val: 1}
	root.Left.Left = &utils.TreeNode{Val: 6}
	root.Left.Right = &utils.TreeNode{Val: 2}
	root.Right.Left = &utils.TreeNode{Val: 0}
	root.Right.Right = &utils.TreeNode{Val: 8}
	p := root.Left
	q := root.Left.Right
	lca := LowestCommonAncestor(root, p, q)
	if lca != root.Left {
		t.Fatalf("LCA expected 5 got %v", lca.Val)
	}
}
