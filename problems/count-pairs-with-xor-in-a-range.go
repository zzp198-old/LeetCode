package problems

//	func countPairs(nums []int, low int, high int) (answer int) {
//		for i := 0; i < len(nums); i++ {
//			for j := i + 1; j < len(nums); j++ {
//				xor := nums[i] ^ nums[j]
//				if xor >= low && xor <= high {
//					answer++
//				}
//			}
//		}
//		return
//	}
func countPairs(nums []int, low, high int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for high++; high > 0; high >>= 1 {
		nxt := map[int]int{}
		for x, c := range cnt {
			if high&1 > 0 {
				ans += c * cnt[x^(high-1)]
			}
			if low&1 > 0 {
				ans -= c * cnt[x^(low-1)]
			}
			nxt[x>>1] += c
		}
		cnt = nxt
		low >>= 1
	}
	return ans / 2
}
