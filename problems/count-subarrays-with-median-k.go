package problems

func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	n := len(nums)
	cnt, x := make([]int, n*2), n
	cnt[x] = 1
	for i := pos - 1; i >= 0; i-- { // 从 pos-1 开始累加 x
		if nums[i] < k {
			x++
		} else {
			x--
		}
		cnt[x]++
	}

	ans, x := cnt[n]+cnt[n-1], n
	for _, v := range nums[pos+1:] { // 从 pos+1 开始累加 x
		if v > k {
			x++
		} else {
			x--
		}
		ans += cnt[x] + cnt[x-1]
	}
	return ans
}
