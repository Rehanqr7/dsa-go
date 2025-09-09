package trees

import (
	"testing"

	"github.com/Rehanqr7/dsa-go/utils"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &utils.TreeNode{Val: 1}
	root.Left = &utils.TreeNode{Val: 2}
	root.Right = &utils.TreeNode{Val: 3}
	root.Left.Left = &utils.TreeNode{Val: 4}
	root.Left.Right = &utils.TreeNode{Val: 5}
	if got := DiameterOfBinaryTree(root); got != 3 {
		t.Fatalf("Diameter expected 3 got %d", got)
	}
}
