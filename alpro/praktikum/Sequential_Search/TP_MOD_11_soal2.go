package SequentialSearch

import (
	atribut "Project-Alpro/alpro/praktikum/atributPraktikum"
	"fmt"
)

func SeqSearch2(T atribut.Array, N, X int) int {
	var idx, k int
	idx = 1
	k = 0
	for idx == -1 && k < N {
		if T.Info[k] == X {
			idx = k
		}
		k = k + 1
	}
	return idx
}

func Soal2SeqSearch() {
	fmt.Println(`
============================================
package main

import (
	atribut "Project-Alpro/alpro/praktikum/atributPraktikum"
	"fmt"
)

func SeqSearch2(T atribut.Array, N, X int) int {
	var idx, k int
	idx = 1
	k = 0
	for idx == -1 && k < N {
		if T.Info[k] == X {
			idx = k
		}
		k = k + 1
	}
	return idx
}

func main() {
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	SeqSearch2(data, data.N, x)

}
============================================`)
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	SeqSearch2(data, data.N, x)

}
