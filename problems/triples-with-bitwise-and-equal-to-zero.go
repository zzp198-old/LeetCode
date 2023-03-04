package problems

func countTriplets(nums []int) int {
	cnt := make(map[int]int)
	for i := range nums {
		for j := range nums {
			cnt[nums[i]&nums[j]]++
		}
	}
	res := 0
	for i := range nums {
		for k, v := range cnt {
			if k&nums[i] == 0 {
				res += v
			}
		}
	}
	return res
}
