package Procedure

import "fmt"

var pi float64 = 3.14

func LuasKelilingLingkaran(r float64, LL, KL *float64) {
	*LL = pi * r * r
	*KL = 2 * pi * r
}
func LuasKelilingPersgi(s float64, LP, KP *float64) {
	*LP = s * s
	*KP = 4 * s
}

func total(LP, LL, KP, KL float64, totalluas *float64, totalkeliling *float64) {
	*totalluas = LL + LP
	*totalkeliling = KP + KL
}

func soal1prosedur() {
	var radius, sisi float64
	var LL, KL, LP, KP float64

	for {
		fmt.Scan(&radius, &sisi)

		if radius != 0 || sisi != 0 {
			LuasKelilingLingkaran(radius, &LL, &KL)
			LuasKelilingPersgi(sisi, &LP, &KP)
			var totalluas, totalkeliling float64
			total(LP, LL, KP, KL, &totalluas, &totalkeliling)
			fmt.Printf("%-10.2s %-10.2s %-10.2s %-10.2s %-10.2s %-10.2s %-10.2s %-10.2s\n", "R", "S", "LL", "LP", "KL", "KP", "LT", "TP")
			fmt.Printf("%-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f\n\n\n", radius, sisi, LL, LP, KL, KP, totalluas, totalkeliling)
		} else {
			break
		}
	}
}
