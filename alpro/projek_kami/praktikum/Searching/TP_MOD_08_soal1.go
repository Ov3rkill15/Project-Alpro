package Searching

import "fmt"

func soal1searching() {
	const NMAX int = 10
	var A [NMAX]int
	var i, n, MAX int
	fmt.Scan(&n)
	if n > NMAX {
		n = NMAX
	}
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MAX = 0
	for i = 1; i < n; i++ {
		if A[i] > A[MAX] {
			MAX = i
		}
	}
	fmt.Println(A[MAX], MAX)
}
