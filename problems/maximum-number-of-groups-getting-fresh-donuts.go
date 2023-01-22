package problems

func maxHappyGroups(m int, groups []int) (ans int) {
	cnt := make([]int, m)
	for _, x := range groups {
		x %= m
		if x == 0 {
			ans++ // 直接排在最前面
		} else if cnt[m-x] > 0 {
			cnt[m-x]-- // 配对
			ans++
		} else {
			cnt[x]++
		}
	}

	val, mask := []int{}, 0
	for x, c := range cnt {
		if c > 0 {
			val = append(val, x)
			mask = mask<<5 | c
		}
	}
	// 上面加入 val 的顺序和拼接 mask 的顺序是相反的，reverse 后就对上了
	for i, n := 0, len(val); i < n/2; i++ {
		val[i], val[n-1-i] = val[n-1-i], val[i]
	}

	// 直接用 pair 当作 key，更方便
	type pair struct{ left, mask int }
	cache := map[pair]int{}
	var dfs func(int, int) int
	dfs = func(left, mask int) (res int) {
		p := pair{left, mask}
		if v, ok := cache[p]; ok {
			return v
		}
		for i, x := range val { // 枚举顾客
			i *= 5
			if mask>>i&31 > 0 { // cnt[x] > 0
				r := dfs((left+x)%m, mask-1<<i)
				if left == 0 {
					r++
				}
				res = max(res, r)
			}
		}
		cache[p] = res
		return
	}
	return ans + dfs(0, mask)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
