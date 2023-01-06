package problems

//func countEven(num int) (answer int) {
//	for i := 1; i <= num; i++ {
//		if bitsum(i)%2 == 0 {
//			answer++
//		}
//	}
//	return
//}

//func bitsum(n int) (sum int) {
//	for n != 0 {
//		sum += n % 10
//		n = n / 10
//	}
//	return
//}

func countEven(num int) (answer int) {
	for i := 1; i <= num; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum%2 == 0 {
			answer++
		}
	}
	return
}
