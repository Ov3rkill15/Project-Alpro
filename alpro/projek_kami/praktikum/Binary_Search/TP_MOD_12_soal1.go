package Binsearch

func binSearch1(tab tabInt, n, x int) bool {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && tab[mid] != x {
		if x < tab[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > -1 && tab[mid] == x
}
