package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{1, 2, 2, 3}, 2))
}

//https://leetcode.com/problems/remove-element/description/

func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	nums = nums[:n]
	return n
}
