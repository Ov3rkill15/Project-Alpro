package Binsearch

func binSearch2(tab tabInt, n, x int) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && tab[mid] != x {
		mid = (left + right) / 2
		if x < tab[mid] {
			right = mid - 1
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}
