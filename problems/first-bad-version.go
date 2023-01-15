package problems

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

// 学到了,sort的Search传入搜索条件,自动二分查找
func firstBadVersion(n int) int {
	return sort.Search(n, isBadVersion)
}
