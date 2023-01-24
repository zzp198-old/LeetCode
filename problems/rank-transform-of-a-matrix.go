package problems

import (
	"math"
	"sort"
)

type Node struct {
	Val int
	//W int
	R  int
	C  int
	Fa int //记录元素的家族
}

type Nodes []Node

func (ns Nodes) Len() int {
	return len(ns)
}

func (ns Nodes) Less(i, j int) bool {
	if ns[i].Val != ns[j].Val {
		return ns[i].Val < ns[j].Val
	} else {
		return ns[i].Fa < ns[j].Fa
	}
}

func (ns Nodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

type Node2 struct {
	Val int
	R   int
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	fa := make([][]int, m)
	for i := 0; i < m; i++ {
		fa[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fa[i][j] = i*m + j
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if fa[i][j] != i*m+j {
				continue
			}
			har := make([]bool, m)
			hac := make([]bool, n)
			har[i] = true
			hac[j] = true

			unionr(i, j, i*m+j, har, hac, fa, matrix)
			unionc(i, j, i*m+j, har, hac, fa, matrix)
		}
	}

	// fmt.Println(fa)

	ns := make(Nodes, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ns = append(ns, Node{matrix[i][j], i, j, fa[i][j]})
		}
	}
	sort.Sort(ns)
	rv := make([]Node2, m)
	for i := 0; i < m; i++ {
		rv[i].Val = math.MinInt32
	}
	cv := make([]Node2, n)
	for i := 0; i < n; i++ {
		cv[i].Val = math.MinInt32
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i, node := range ns {
		r := rv[node.R].R
		if node.Val > rv[node.R].Val {
			r++
		}
		if node.Val == cv[node.C].Val && cv[node.C].R > r {
			r = cv[node.C].R
		}
		if node.Val > cv[node.C].Val && cv[node.C].R >= r {
			r = cv[node.C].R + 1
		}
		result[node.R][node.C] = r
		rv[node.R].Val, rv[node.R].R = node.Val, r
		cv[node.C].Val, cv[node.C].R = node.Val, r

		for j := i - 1; j >= 0; j-- {
			now := ns[j]
			if node.Fa != now.Fa || result[node.R][node.C] <= result[now.R][now.C] {
				break
			}
			result[now.R][now.C] = result[node.R][node.C]
			rv[now.R].R, cv[now.C].R = result[node.R][node.C], result[node.R][node.C]

		}
	}

	return result
}

// 合并同一行的相同值，并会尝试合并元素所在的列的相同值
func unionr(r, c, root int, har, hac []bool, fa [][]int, matrix [][]int) {
	_, n := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		if matrix[r][i] == matrix[r][c] && fa[r][i] != root {
			fa[r][i] = root
			if !hac[i] {
				hac[i] = true
				unionc(r, i, root, har, hac, fa, matrix)
			}
		}
	}
}

// 合并同一列的相同值，并会尝试合并元素所在的行的相同值
func unionc(r, c, root int, har, hac []bool, fa [][]int, matrix [][]int) {
	// fmt.Println("bbb", r, c)
	m, _ := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][c] == matrix[r][c] && fa[i][c] != root {
			fa[i][c] = root
			// fmt.Println("bbb", r, c, i)
			if !har[i] {
				har[i] = true
				unionr(i, c, root, har, hac, fa, matrix)
			}
		}
	}
}
