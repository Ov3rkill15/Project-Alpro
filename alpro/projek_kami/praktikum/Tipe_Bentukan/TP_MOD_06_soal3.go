package Tipebentukan

import "fmt"

// Tipe bentukan pegawai dengan field nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama                   string
	gaji_pokok, masa_kerja int
	besar_bonus            float64
}

func Soal3bentukan() {
	// deklarasi variabel bertipe pegawai
	var names pegawai
	// baca data pengawai
	fmt.Scan(&names.nama, &names.gaji_pokok, &names.masa_kerja)
	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&names)
	// Cetak besar bonus
	fmt.Printf("pegawai dengan nama %s mendapatkan bonus sebesar %.f", names.nama, names.besar_bonus)
}
func hitung_bonus(p *pegawai) {
	var angka_pengali float64
	if p.masa_kerja >= 10 {
		angka_pengali = 1.5
	} else if p.masa_kerja >= 5 {
		angka_pengali = 0.75
	} else {
		angka_pengali = 0.5
	}
	p.besar_bonus = float64(p.gaji_pokok) * angka_pengali

}
