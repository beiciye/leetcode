package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maxSubArray([]int{8, -19, 5, -4, 20}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := nums[0]
	lastSum := 0
	gapSum := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if sum < nums[i] {

			}

			if (gapSum + nums[i]) > 0 {
				sum = sum + gapSum + nums[i]
				gapSum = 0
				lastSum = 0
			}
			// } else {
			// 	if lastSum > 0 {
			// 		if nums[i]+lastSum > sum {
			// 			sum = nums[i] + lastSum
			// 			lastSum = 0
			// 			gapSum = 0
			// 		} else {
			// 			gapSum += nums[i]
			// 			lastSum += nums[i]
			// 		}
			// 	} else {
			// 		if nums[i] > sum {
			// 			sum = nums[i]
			// 			lastSum = 0
			// 			gapSum = 0
			// 		} else {
			// 			lastSum = nums[i]
			// 			gapSum += nums[i]
			// 		}
			// 	}
		} else {
			if sum < nums[i] {
				sum = nums[i]
				lastSum = 0
				gapSum = 0
			} else {
				gapSum += nums[i]
				if lastSum > 0 {
					lastSum = nums[i] + lastSum
				} else {
					lastSum = nums[i]
				}

			}

		}
	}
	fmt.Println("lastnum", lastSum)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
