package trees

import "github.com/Rehanqr7/dsa-go/utils"

/*
Binary Tree Traversals (Easy)
- Inorder, Preorder, Postorder (DFS)
- Level Order (BFS)

Time Complexity: O(n)
Space Complexity: O(n) for recursion/queue
*/

func Inorder(root *utils.TreeNode) []int {
	res := []int{}
	var dfs func(*utils.TreeNode)
	dfs = func(n *utils.TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		res = append(res, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

func Preorder(root *utils.TreeNode) []int {
	res := []int{}
	var dfs func(*utils.TreeNode)
	dfs = func(n *utils.TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

func Postorder(root *utils.TreeNode) []int {
	res := []int{}
	var dfs func(*utils.TreeNode)
	dfs = func(n *utils.TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		dfs(n.Right)
		res = append(res, n.Val)
	}
	dfs(root)
	return res
}

func LevelOrder(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	q := []*utils.TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			level = append(level, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
