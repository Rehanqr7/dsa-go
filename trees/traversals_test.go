package trees

import (
	"reflect"
	"testing"

	"github.com/Rehanqr7/dsa-go/utils"
)

func TestTraversals(t *testing.T) {
	root := &utils.TreeNode{Val: 1,
		Left: &utils.TreeNode{Val: 2,
			Left:  &utils.TreeNode{Val: 4},
			Right: &utils.TreeNode{Val: 5},
		},
		Right: &utils.TreeNode{Val: 3},
	}
	if got := Inorder(root); !reflect.DeepEqual(got, []int{4, 2, 5, 1, 3}) {
		t.Fatalf("Inorder got %v", got)
	}
	if got := Preorder(root); !reflect.DeepEqual(got, []int{1, 2, 4, 5, 3}) {
		t.Fatalf("Preorder got %v", got)
	}
	if got := Postorder(root); !reflect.DeepEqual(got, []int{4, 5, 2, 3, 1}) {
		t.Fatalf("Postorder got %v", got)
	}
	lvl := LevelOrder(root)
	exp := [][]int{{1}, {2, 3}, {4, 5}}
	if !reflect.DeepEqual(lvl, exp) {
		t.Fatalf("LevelOrder got %v want %v", lvl, exp)
	}
}
