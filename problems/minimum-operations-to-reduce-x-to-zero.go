package problems

import "math"

// 取最长子段,合为sum-target.长度固定为滑动窗口，不固定为双指针
func minOperations(nums []int, x int) int {
	target := -x
	for _, num := range nums {
		target += num
	}
	if target < 0 { // 全部移除也不够
		return -1
	}
	answer, left, sum := -1, 0, 0
	for right, num := range nums { // 右边界一直右移，左边界看情况右移
		sum += num
		for sum > target { // 减小子段长度,left右移
			sum -= nums[left]
			left++
		}
		if sum == target {
			answer = int(math.Max(float64(answer), float64(right-left+1)))
		}
	}
	if answer < 0 {
		return -1
	}
	return len(nums) - answer
}
