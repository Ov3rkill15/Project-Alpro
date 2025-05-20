package Searching

import "fmt"

const NMAX = 20

type tabInt [NMAX]int

func soal3searching() {
	var A tabInt
	var n int

	baca(&A, &n)
	cetakElemen(A, n)
	cetakInfo(A, n)
}

func baca(A *tabInt, n *int) {
	var num int
	var stop bool
	*n = 0
	for *n < NMAX && !stop {
		fmt.Scan(&num)
		if num > 0 {
			A[*n] = num
			*n++
		} else {
			stop = true
		}
	}
}

func cetakElemen(A tabInt, n int) {
	fmt.Print("Elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	var max, i int
	if n == 0 {
		return 0
	}
	max = A[0]
	for i = 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	var min, i int
	if n == 0 {
		return 0
	}
	min = A[0]
	for i = 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	fmt.Printf("Banyak elemen: %d\n", n)
}
