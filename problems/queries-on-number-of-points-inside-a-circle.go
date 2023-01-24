package problems

func countPoints(points [][]int, queries [][]int) []int {
	answer := make([]int, len(queries))
	for _, point := range points {
		x := point[0]
		y := point[1]
		for i, query := range queries {
			l2 := (x-query[0])*(x-query[0]) + (y-query[1])*(y-query[1])
			if l2 <= query[2]*query[2] {
				answer[i]++
			}
		}
	}
	return answer
}
