package SequentialSearch

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func Soal3SeqSearch() {
	var data tabInt
	var nData int
	var x int

	fmt.Scan(&x)
	fmt.Scan(&nData)

	if nData > NMAX {
		nData = NMAX
	}
	bacaData(&data, nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x) {
		freq := frekuensiBilangan(data, nData, x)
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", freq, x)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}
func bacaData(A *tabInt, n int) {
	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		A[i] = num
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", A[i])
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	count := 0
	for i := 0; i < n; i++ {
		if A[i] == x {
			count++
		}
	}
	return count
}

func sequentialSearch(A tabInt, n, x int) bool {
	for i := 0; i < n; i++ {
		if A[i] == x {
			return true
		}
	}
	return false
}
