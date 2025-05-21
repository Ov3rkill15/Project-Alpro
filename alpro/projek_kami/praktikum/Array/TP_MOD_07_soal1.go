package Array

import "fmt"

func Soal1array() {
	var n int
	const NMAX = 10
	fmt.Println("Masukkan jumlah data: (maks 10)")
	fmt.Scan(&n)
	var A [NMAX]string
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println()
	for i = 0; i < n; i++ {
		fmt.Println(A[i])
	}
}
