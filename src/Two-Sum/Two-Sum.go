package twoSum

// https://leetcode.com/problems/two-sum/description/

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
//    Given nums = [2, 7, 11, 15], target = 9,
//    Because nums[0] + nums[1] = 2 + 7 = 9,
//    return [0, 1].

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int, 0)
	for idx, val := range nums {
		temp[val] = idx
	}
	for idx, ele := range nums {
		needed := target - ele
		if index, ok := temp[needed]; ok {
			if index == idx {
				continue
			}
			return []int{idx, temp[needed]}
		}
	}
	return []int{0, 0}
}
