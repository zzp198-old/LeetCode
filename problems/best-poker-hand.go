package problems

import "bytes"

//	func bestHand(ranks []int, suits []byte) string {
//		sort.Ints(ranks)
//		//if suits[0] == suits[1] &&
//		//	suits[1] == suits[2] &&
//		//	suits[2] == suits[3] &&
//		//	suits[3] == suits[4] {
//		//	return "Flush"
//		//}
//		if bytes.Count(suits, suits[:1]) == 5 {
//			return "Flush"
//		}
//		if ranks[0] == ranks[2] ||
//			ranks[1] == ranks[3] ||
//			ranks[2] == ranks[4] {
//			return "Three of a Kind"
//		}
//		if ranks[0] == ranks[1] ||
//			ranks[1] == ranks[2] ||
//			ranks[2] == ranks[3] ||
//			ranks[3] == ranks[4] {
//			return "Pair"
//		}
//		return "High Card"
//	}
func bestHand(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	for _, r := range ranks {
		cnt[r]++
		if cnt[r] == 3 {
			return "Three of a Kind"
		}
		if cnt[r] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}
