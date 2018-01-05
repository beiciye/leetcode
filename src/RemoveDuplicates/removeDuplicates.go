package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	i := 0
	for k := 0; k < len(nums); k++ {
		if k == 0 {
			i++
			continue
		}
		if nums[k] != nums[k-1] {
			nums[i] = nums[k]
			i++
		}
	}
	nums = nums[:i]
	fmt.Println(nums)
	return len(nums)
}
