package problems

func minMaxGame(nums []int) int {
	for len(nums) != 1 {
		arr := make([]int, len(nums)/2)
		for i := 0; i < len(nums)/2; i++ {
			if i%2 == 0 {
				arr[i] = min(nums[2*i], nums[2*i+1])
			} else {
				arr[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = arr
	}
	return nums[0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
