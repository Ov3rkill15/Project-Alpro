package Array

import "fmt"

type tabInt struct {
	info [NMAX]int
	n    int
}

func bacaData(A *tabInt) {
	var x int
	if A.n < NMAX {
		fmt.Scan(&x)
		A.info[A.n] = x
		A.n++
	}
}

func cetakData(A tabInt) {
	if A.n > 0 {
		for i := 0; i < A.n; i++ {
			fmt.Print(A.info[i], " ")
		}
		fmt.Println()
	}
}

func soal3array() {
	var data tabInt
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	cetakData(data)
}
