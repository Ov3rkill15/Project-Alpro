package Fungsi

import (
	"fmt"
	"unicode"
)

func lowToUpper(char rune) rune {
	return unicode.ToUpper(char)
}
func soal2fungsi() {
	var input rune
	fmt.Scanf("%c", &input)
	fmt.Printf("%c\n", lowToUpper(input))

}
