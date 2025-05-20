package Array

import (
	atribut "Project-Alpro/alpro/projek_kami/praktikum/atributPraktikum"
	"fmt"
)

func soal2array() {
	var data atribut.Tabint
	var banyakData int
	atribut.BacaData(&data, &banyakData)
	fmt.Println()
	atribut.CetakData(data, banyakData)

}
