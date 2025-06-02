package Array

import "fmt"

const NMAX int = 10

type Tabint [NMAX]int

func Soal3array() {
	fmt.Println(`
============================================
package main

import "fmt"

const NMAX int = 10

type Tabint [NMAX]int

func main() {
	var data Tabint
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData))
	fmt.Println(rataRata(data, nData))
}

func baca(A *Tabint, n *int) {
	var num int
	var stop bool
	*n = 0
	stop = true

	for stop && *n < NMAX {
		fmt.Scan(&num)

		if num == 0 {
			stop = false
		} else {
			if num < 0 {
				num = -num
			}
			A[*n] = num
			*n++
		}
	}
}

func cetak(A Tabint, n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Print(A[i], " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Array kosong")
	}
}

func jumlah(A Tabint, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	return total
}

func rataRata(A Tabint, n int) float64 {
	if n > 0 {
		return float64(jumlah(A, n)) / float64(n)
	}
	return 0.0
}

============================================
			`)

	var data Tabint
	var nData int
	fmt.Println("SILAHKAN DICOBA: ")
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData))
	fmt.Println(rataRata(data, nData))
}

func baca(A *Tabint, n *int) {
	var num int
	var stop bool
	*n = 0
	stop = true

	for stop && *n < NMAX {
		fmt.Scan(&num)

		if num == 0 {
			stop = false
		} else {
			if num < 0 {
				num = -num
			}
			A[*n] = num
			*n++
		}
	}
}

func cetak(A Tabint, n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Print(A[i], " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Array kosong")
	}
}

func jumlah(A Tabint, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	return total
}

func rataRata(A Tabint, n int) float64 {
	if n > 0 {
		return float64(jumlah(A, n)) / float64(n)
	}
	return 0.0
}
