package problems

//func digitCount(num string) bool {
//	bit := make(map[int]int, 10) // 存储字符串的数字统计
//	for _, a := range num {
//		i, _ := strconv.Atoi(string(a))
//		bit[i]++
//	}
//	for i := 0; i < len(num); i++ {
//		j, _ := strconv.Atoi(string(num[i]))
//		if j != bit[i] {
//			return false
//		}
//	}
//	return true
//}

func digitCount(num string) bool {
	cnt := map[rune]int{}
	for _, c := range num {
		cnt[c-'0']++
	}
	for i, c := range num {
		if cnt[rune(i)] != int(c-'0') {
			return false
		}
	}
	return true
}
