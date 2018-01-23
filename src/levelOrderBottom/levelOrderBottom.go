package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	s := make([][]int, 0)
	tree(&s, root, 0)
	return s
}

func tree(s *[][]int, t *TreeNode, deep int) {
	if t == nil {
		return
	}
	if len(*s) == deep {
		*s = append(*s, []int{})
	}

	(*s)[deep] = append((*s)[deep], t.Val)

	if t.Left != nil {
		tree(s, t.Left, deep+1)
	}
	if t.Right != nil {
		tree(s, t.Right, deep+1)
	}
}
