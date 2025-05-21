package Fungsi

import "fmt"

func check(bilangan_input, bilangan_cek int) bool {
	var temp int
	var found bool
	for bilangan_input > 0 {
		temp = bilangan_input % 10
		if temp == bilangan_cek {
			found = true
			break
		}
		bilangan_input /= 10
	}
	return found
}

func Soal3fungsi() {
	var b1, b2 int
	var found bool

	fmt.Scan(&b1, &b2)
	found = check(b1, b2)
	if found {
		fmt.Println("YA")
	} else {
		fmt.Println("TIDAK")
	}
}
