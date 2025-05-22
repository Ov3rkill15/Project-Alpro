package Binsearch

import (
	atribut "Project-Alpro/alpro/projek_kami/praktikum/atributPraktikum"
	"fmt"
)

func BinSearch2(tab atribut.Array, n, x int) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && tab.Info[mid] != x {
		mid = (left + right) / 2
		if x < tab.Info[mid] {
			right = mid - 1
		} else if x > tab.Info[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func Soal2BinSearch() {
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	BinSearch2(data, data.N, x)

}
