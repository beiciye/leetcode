package main

import (
	"bytes"
	"fmt"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	p1 := &bytes.Buffer{}
	q1 := &bytes.Buffer{}

	tree(p1, p, 0)
	tree(q1, q, 0)
	p1s := p1.String()
	q1s := q1.String()

	return p1s == q1s
}

func tree(buf *bytes.Buffer, t *TreeNode, deep int) {
	if t == nil {
		return
	}
	buf.WriteString(fmt.Sprintf("%d:%d", deep, t.Val))

	if t.Left != nil {
		buf.WriteString("l")
		tree(buf, t.Left, deep+1)

	}
	if t.Right != nil {
		buf.WriteString("r")
		tree(buf, t.Right, deep+1)
	}
}
