package problems

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Sort(sort.IntSlice(coins))
	if coins[0] > 1 {
		return 1
	}
	ma := 0
	n := len(coins)
	for i := 0; i < n; i++ {
		if ma >= coins[i]-1 {
			ma += coins[i]
		} else {
			break
		}
	}
	return ma + 1
}
