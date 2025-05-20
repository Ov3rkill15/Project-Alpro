package Array

import "fmt"

func soal1array() {
	var n int
	const NMAX = 10
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
