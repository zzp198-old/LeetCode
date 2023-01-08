package problems

//	func mySqrt(x int) int {
//		for i := 0; i < math.MaxInt; i++ {
//			if i*i <= x {
//				continue
//			} else {
//				return i - 1
//			}
//		}
//		return -1
//	}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	r := x
	for r > x/r {
		r = int((r + x/r) / 2)
	}
	return r
}
