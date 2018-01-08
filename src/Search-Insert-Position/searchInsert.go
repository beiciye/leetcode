package main

func main() {

}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	for idx, val := range nums {
		if val >= target {
			return idx
		}
	}
	return -1
}
