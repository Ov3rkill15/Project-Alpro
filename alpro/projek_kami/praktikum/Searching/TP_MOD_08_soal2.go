package Searching

import "fmt"

func soal2searching() {
	const NMAX int = 10
	var A [NMAX]int
	var i, n, MIN int
	fmt.Scan(&n)
	if n > NMAX {
		n = NMAX
	}
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MIN = 0
	for i = 1; i < n; i++ {
		if A[i] < A[MIN] {
			MIN = i
		}
	}
	fmt.Println(A[MIN], MIN)
}
