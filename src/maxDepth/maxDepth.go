package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return tree(root, 0)
}

func tree(t *TreeNode, i int) int {
	if t == nil {
		return i
	}
	return max(tree(t.Left, i+1), tree(t.Right, i+1))
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
