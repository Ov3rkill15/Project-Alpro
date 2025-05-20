package atributpraktikum

import "fmt"

const NMAX int = 1000

type Tabint [NMAX]int

func BacaData(A *Tabint, N *int) {
	var i int
	for i = 0; i < *N; i++ {
		fmt.Scan(&A[i])
	}
}

func CetakData(A Tabint, N int) {
	for i := 0; i < N; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
