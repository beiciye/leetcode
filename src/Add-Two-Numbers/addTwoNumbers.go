package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/description/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 9,
	}
	l2 := &ListNode{
		Val: 1,
	}
	q := addTwoNumbers(l1, l2)

	fmt.Println(q, q.Next)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i := 0
	temp := &ListNode{}
	result := &ListNode{Next: temp}
	for l1 != nil || l2 != nil {
		val := l1.Val + l2.Val + i
		upper := false
		if val >= 10 {
			val = val - 10
			upper = true
		}
		temp.Val = val
		temp.Next = &ListNode{}
		i = 0
		if upper {
			i = 1
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && i == 0 {
			temp.Next = nil
			break
		} else {
			temp = temp.Next
		}

		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
	}
	return result.Next
}
