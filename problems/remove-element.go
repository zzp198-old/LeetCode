package problems

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			nums[l] = nums[r-1]
			r--
		} else {
			l++
		}
	}
	return l
}
