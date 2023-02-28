package problems

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hash := map[int]int{}
	for _, item := range items1 {
		hash[item[0]] += item[1]
	}
	for _, item := range items2 {
		hash[item[0]] += item[1]
	}
	var ret [][]int
	for k, v := range hash {
		ret = append(ret, []int{k, v})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i][0] < ret[j][0]
	})
	return ret
}
