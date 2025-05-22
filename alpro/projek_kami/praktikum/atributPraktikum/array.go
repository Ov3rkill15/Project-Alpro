package atributpraktikum

import "fmt"

const NMAX int = 1000

type Array struct {
	Info [NMAX]int
	N    int
}

func BacaData1(A *Array) {
	fmt.Scan(&A.N)
	if A.N > NMAX {
		A.N = NMAX
	}
	for i := 0; i < A.N; i++ {
		fmt.Scan(&A.Info[i])
	}
}

func CetakData1(A Array) {
	if A.N == 0 {
		fmt.Println("Array kosong")
	}
	for i := 0; i < A.N; i++ {
		fmt.Print(A.Info[i], " ")
	}
	fmt.Println()
}
