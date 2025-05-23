package SequentialSearch

import (
	atribut "Project-Alpro/alpro/projek_kami/praktikum/atributPraktikum"
	"fmt"
)

func SeqSearch(T atribut.Array, N, X int) bool {
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for !ketemu && k < N {
		ketemu = T.Info[k] == X
		k = k + 1
	}
	return ketemu
}

func Soal1SeqSearch() {
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	SeqSearch(data, data.N, x)

}
