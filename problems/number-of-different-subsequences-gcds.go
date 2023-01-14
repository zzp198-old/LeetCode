package problems

func countDifferentSubsequenceGCDs(nums []int) (answer int) {
	mx := 0
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}
	has := make([]bool, mx+1)
	for _, x := range nums {
		if !has[x] {
			has[x] = true
			answer++
		}
	}
	for i := 1; i <= mx/3; i++ {
		if has[i] {
			continue
		}
		g := 0
		for j := i * 2; j <= mx && g != i; j += i {
			if has[j] {
				g = gcd(g, j)
			}
		}
		if g == i {
			answer++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
