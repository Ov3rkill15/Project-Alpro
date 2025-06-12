package main

import (
	"Project-Alpro/alpro"
	"Project-Alpro/atribut"
	"Project-Alpro/login"
	"Project-Alpro/pengpro"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	atribut.ClearScreen()
	reader := bufio.NewReader(os.Stdin)
	var stop bool = false // Variabel untuk menghentikan loop
	var sign, helo string
	login.Mainlogin(&sign, &helo)
	switch helo {
	case "nathasyayuanmaharani":
		helo = "Yuan"
	case "theodoreelvisestrada":
		helo = "Elvis"
	case "dyahkusumawardani":
		helo = "Dyah"
	case "azrielraihaneldovahartoto":
		helo = "Azriel Jids"
	case "muhammadilhamalifianda":
		helo = "Ilham"
	case "alyaazizaputeri":
		helo = "Alya"
	case "ahmadabdansyakuro":
		helo = "Abdan"
	case "fathurrahmanalfarizi":
		helo = "Fathur"
	case "nuswantorosetyomukti":
		helo = "Nuswantoro"
	case "anggitacahyatihidayat":
		helo = "Anggita"
	case "wibnuhijrahfranstio":
		helo = "Wibnu"
	case "meyshaprimiandita":
		helo = "Meysha"
	case "muhamadfiqrihabibi":
		helo = "Fiqri"
	case "fitriacahyani":
		helo = "Fitria"
	case "triansyahdaniswaraibrahim":
		helo = "Triansyah"
	case "rakhaabdillahalkautsar":
		helo = "Rakha"
	case "avicenanaufallathif":
		helo = "Avicenna"
	case "naylaassyifa":
		helo = "Nayla"
	case "williampetervanxnajoan":
		helo = "William"
	case "rayvanalifarlomahesworo":
		helo = "Rayvan"
	case "zaidansalamrojab":
		helo = "Zaidan"
	case "audreyfredileyhanas":
		helo = "Audrey"
	case "muhammadnaelfadly":
		helo = "Nael"
	case "nairacahayaputridarmawansinaga":
		helo = "Naira"
	case "muhamadalwansuryadi":
		helo = "Alwan"
	case "dhafyahmadzubaidi":
		helo = "Dhafy"
	case "muhammadfarisdhiyaylhaqsarbini":
		helo = "Faris"
	case "nursyadira":
		helo = "Nora"
	case "rayfitokrisnawijaya":
		helo = "Fito"
	case "mochammadrafirisqullah":
		helo = "Rafi"
	case "iputugedeagastyakrisnawidartha":
		helo = "Agastya"
	case "rendil":
		helo = "Rendil"
	case "muhammadariqazzaki":
		helo = "Ariq"
	case "edmundyuliusgantur":
		helo = "Edmund"
	case "muhammadsayyidhuwaidi":
		helo = "Sayyit kumis"
	case "muhdzuljalalwaliikramjalil":
		helo = "Dzul"
	case "ramadhantangguhdefennder":
		helo = "SENSEI"
	case "adzkiyaputrirahmawan":
		helo = "Adzkiya"
	case "fathimahradhiyya":
		helo = "Radhi"
	case "rakanghazianadiwjaya":
		helo = "Rakan"
	case "jihannabilamubarakah":
		helo = "Jihan"
	}

	for !stop {
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan menu
		user := figure.NewFigure(helo, "doom", true).String()
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		fmt.Print("\033[32m")
		fmt.Print(user)
		fmt.Print("\033[0m")
		menu()
		fmt.Println(atribut.GetWaktuSekarang())
		fmt.Print("Masukkan pilihan: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			atribut.ClearScreen()
			fmt.Println("Menuju Pengenalan Pemrograman...")
			atribut.Loading(1200)
			pengpro.MainMenu()
		case "2":
			atribut.ClearScreen()
			fmt.Println("Menuju Algoritma Pemrograman...")
			atribut.Loading(1200)
			alpro.MainMenu()
		case "3":
			atribut.ClearScreen()
			main()
		case "10":
			atribut.ClearScreen()
			atribut.BukaPDF("D:\\Matkul smester 2\\Project-Alpro\\cv\\CV_Muhamad Alwan Suryadi.pdf")
			atribut.BukaPDF("D:\\Matkul smester 2\\Project-Alpro\\cv\\CV__103032400035_fathurrahman alfarizi.pdf")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading(1200)
		case "11":
			atribut.ClearScreen()
			atribut.BukaPDF("D:\\Matkul smester 2\\Project-Alpro\\cv\\LAPORAN_TUBES_MUHAMMAD_ALWAN_SURYADI & FATHURRAHMAN_ALFARIZI_KELOMPOK_18.pdf")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading(1200)
		case "0":
			stop = true
			atribut.ClearScreen()
			figure.NewFigure("THX THX THX", "basic", true).Blink(5000, 1000, 500)

		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(1200)
			fmt.Println()
		}

	}

}

func menu() {
	letsgo := figure.NewFigure("LET'S GO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(letsgo)     // Cetak teks ASCII
	fmt.Print("\033[0m")  // Reset warna ke default
	fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Materi yang tersedia:
=========================================
1. Pengenalan Pemrograman
2. Algoritma dan Pemrograman
3. Logout


=========================================
10. about us:
11. about this app:
=========================================

0. keluar
	`)
}
