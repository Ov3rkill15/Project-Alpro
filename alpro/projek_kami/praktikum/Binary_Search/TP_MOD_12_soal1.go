package Binsearch

import (
	atribut "Project-Alpro/alpro/projek_kami/praktikum/atributPraktikum"
	"fmt"
)

func BinSearch1(tab atribut.Array, n, x int) bool {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && tab.Info[mid] != x {
		if x < tab.Info[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > -1 && tab.Info[mid] == x
}

func Soal1BinSearch() {
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	BinSearch1(data, data.N, x)

}
