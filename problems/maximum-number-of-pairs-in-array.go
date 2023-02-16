package problems

func numberOfPairs(nums []int) []int {
	pair := 0
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
		if hash[num] == 2 {
			hash[num] = 0
			pair++
		}
	}
	return []int{pair, len(nums) - 2*pair}
}
