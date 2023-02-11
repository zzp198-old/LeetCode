package problems

import "strings"

func alphabetBoardPath(target string) string {
	str := strings.Builder{}
	cx, cy, tx, ty := 0, 0, 0, 0
	for _, char := range target {
		tx = int(char-'a') % 5
		ty = int(char-'a') / 5
		for { // Z处不能右走,不能直下走左到Z
			if cy > ty { // 上
				cy--
				str.WriteByte('U')
				continue // 一直走,而不是来回打折,优化可以跳过N次循环
			}
			if cx > tx { // 在当前左边
				cx--
				str.WriteByte('L')
				continue
			}
			if cy < ty { // 在当前下边
				cy++
				str.WriteByte('D')
				continue
			}
			if cx < tx { // 在当前右边
				cx++
				str.WriteByte('R')
				continue
			}
			if cx == tx && cy == ty {
				break
			}
		}
		str.WriteByte('!')
	}
	return str.String()
}
