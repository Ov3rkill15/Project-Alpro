package login

import (
	"Project-Alpro/atribut"
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"golang.org/x/term"
)

type User struct {
	Username, Password string
}

const NMAX int = 45

type Daftarnama [NMAX]User

var users Daftarnama
var jumlahPengguna int // Inisialisasi default ke 0

// init() adalah fungsi khusus di Go yang otomatis dieksekusi satu kali
// saat package ini diinisialisasi. Ini adalah tempat yang tepat untuk
// memanggil initUsers() agar data pengguna hanya dimuat sekali.
func init() {
	initUsers()
}

func initUsers() {
	// Data awal pengguna
	data := [42]User{
		{"admin", "admin123"},
		{"nathasyayuanmaharani", "0001"},
		{"theodoreelvisestrada", "0006"},
		{"dyahkusumawardani", "0009"},
		{"azrielraihaneldovahartoto", "0010"},
		{"muhammadilhamalifianda", "0022"},
		{"alyaazizaputeri", "0026"},
		{"ahmadabdansyakuro", "0029"},
		{"fathurrahmanalfarizi", "0035"},
		{"nuswantorosetyomukti", "0040"},
		{"anggitacahyatihidayat", "0041"},
		{"wibnuhijrahfranstio", "0048"},
		{"meyshaprimiandita", "0050"},
		{"muhamadfiqrihabibi", "0056"},
		{"fitriacahyani", "0060"},
		{"triansyahdaniswaraibrahim", "0062"},
		{"rakhaabdillahalkautsar", "0068"},
		{"avicenanaufallathif", "0073"},
		{"naylaassyifa", "0078"},
		{"williampetervanxnajoan", "0084"},
		{"rayvanalifarlomahesworo", "0087"},
		{"zaidansalamrojab", "0088"},
		{"audreyfredileyhanas", "0093"},
		{"muhammadnaelfadly", "0096"},
		{"nairacahayaputridarmawansinaga", "0100"},
		{"muhamadalwansuryadi", "0104"},
		{"dhafyahmadzubaidi", "0107"},
		{"muhammadfarisdhiyaylhaqsarbini", "0117"},
		{"nursyadira", "0123"},
		{"rayfitokrisnawijaya", "0124"},
		{"mochammadrafirisqullah", "0129"},
		{"iputugedeagastyakrisnawidartha", "0134"},
		{"rendil", "0137"},
		{"muhammadariqazzaki", "0138"},
		{"edmundyuliusgantur", "0155"},
		{"muhammadsayyidhuwaidi", "0157"},
		{"muhdzuljalalwaliikramjalil", "0160"},
		{"ramadhantangguhdefennder", "0003"},
		{"adzkiyaputrirahmawan", "0025"},
		{"fathimahradhiyya", "0029"},
		{"rakanghazianadiwjaya", "0034"},
		{"jihannabilamubarakah", "0037"},
	}

	// Jangan mereset jumlahPengguna jika sudah ada pengguna dari sesi sebelumnya (jika ada data persisten)
	// Namun, karena belum ada persistensi ke file, ini akan selalu memulai dari data hardcode.
	jumlahPengguna = 0
	i := 0
	for i < NMAX {
		if i < len(data) {
			users[i] = data[i]
			jumlahPengguna++
		}
		i++
	}
}

// Login sekarang juga mengembalikan boolean kedua untuk menandakan apakah itu logout admin
func Login(sign, helo *string) (bool, bool) { // <-- Perubahan di sini: tambahkan return bool kedua
	var username, password string
	var signup User
	stop := true
	key := 0
	reader := bufio.NewReader(os.Stdin)
	welcome := figure.NewFigure("=== LOGIN ===", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")
	fmt.Println()
	fmt.Println("1. Masuk")
	fmt.Println("2. Daftar")
	fmt.Print("pilih yang mana: ")
	signed, _ := reader.ReadString('\n')
	*sign = strings.TrimSpace(signed)
	cek := strings.ToLower(*sign)

	if cek == "masuk" || cek == "1" {
		welcomeLogin := figure.NewFigure("=== MASUK ===", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcomeLogin)
		fmt.Print("\033[0m")
		for stop && key < 3 {
			fmt.Print("Masukkan Username: ")
			usernameInput, _ := reader.ReadString('\n')
			username = strings.TrimSpace(usernameInput)

			fmt.Print("Masukkan Password: ")
			passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca kata sandi:", err)
				return false, false // Keluar dari fungsi Login, bukan logout admin
			}
			password = string(passwordBytes)

			if AuthenticateUser(username, password) {
				if helo != nil {
					*helo = username
				}
				atribut.ClearScreen()
				fmt.Println("Login berhasil!")
				fmt.Println("Menuju Aplikasi")
				atribut.Loading(1200)
				fmt.Println()
				stop = false // Berhasil login, hentikan loop

				// Perbaikan logika admin
				if *helo == "admin" {
					var adminLogout bool
					MainAdmin("admin", &adminLogout)
					if adminLogout {
						// Jika admin logout, kembalikan false untuk login gagal dan true untuk admin logout
						return false, true // <-- Perubahan di sini: kembalikan true untuk adminLogout
					}
				}
				return true, false // Login berhasil untuk user biasa atau admin yang tidak logout
			} else {
				fmt.Println("Username atau password salah.")
				key++
				fmt.Printf("Tersisa %d kesempatan\n", 3-key)
			}
		}
		return false, false // Login gagal setelah 3 kali percobaan
	} else if cek == "daftar" || cek == "2" {
		welcomeDaftar := figure.NewFigure("=== DAFTAR ===", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcomeDaftar)
		fmt.Print("\033[0m")

		if jumlahPengguna >= NMAX {
			fmt.Println("Maaf, jumlah maksimum pengguna telah tercapai. Tidak bisa Daftar sekarang.")
			return false, false // Keluar dari fungsi jika penuh
		}

		fmt.Print("Masukkan Username: ")
		signup.Username, _ = reader.ReadString('\n')
		signup.Username = strings.TrimSpace(signup.Username)

		// Cek apakah username sudah ada
		if isUsernameExists(signup.Username) {
			fmt.Println("Username sudah digunakan. Silakan pilih username lain.")
			return false, false // Keluar dari fungsi jika username sudah ada
		}

		fmt.Print("Masukkan NIM: ")
		NIM, _ := reader.ReadString('\n') // NIM disimpan tapi tidak digunakan di struct User
		NIM = strings.TrimSpace(NIM)

		fmt.Println("Selamat datang,")
		fmt.Println("Nama:", signup.Username)
		fmt.Println("NIM:", NIM)

		stop = true // Reset stop untuk loop password
		key = 0     // Reset key untuk loop password
		for stop && key < 3 {
			fmt.Print("Masukkan Password: ")
			passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca kata sandi:", err)
				return false, false
			}

			fmt.Print("Konfirmasi Password: ")
			confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca konfirmasi password:", err)
				return false, false
			}

			if string(passwordBytes) != string(confirmBytes) {
				fmt.Println("Password dan konfirmasi tidak cocok.")
				key++
			} else {
				stop = false // Password cocok, hentikan loop
				signup.Password = string(passwordBytes)
				users[jumlahPengguna] = signup
				jumlahPengguna++
				fmt.Println("Signup berhasil! Silakan login dengan akun yang baru dibuat.")
				return false, false // Setelah daftar, kembali ke menu login
			}
		}
		return false, false // Gagal daftar setelah 3 percobaan
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		fmt.Println("\nTekan Enter untuk kembali...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return false, false
	}
}

func AuthenticateUser(username, password string) bool {
	i := 0
	found := false
	for i < jumlahPengguna {
		if users[i].Username == username && users[i].Password == password {
			found = true
		}
		i++
	}
	return found
}

// isUsernameExists untuk mengecek apakah username sudah ada
func isUsernameExists(username string) bool {
	i := 0
	exists := false
	for i < jumlahPengguna {
		if users[i].Username == username {
			exists = true
			// Tidak menggunakan break, jadi loop akan terus berjalan
			// hingga selesai, tapi 'exists' sudah true.
		}
		i++
	}
	return exists
}

func Mainlogin(sign, helo *string) {
	// Loop utama untuk menu login/daftar
	running := true
	for running {
		*helo = ""                                       // Reset helo sebelum setiap percobaan login/daftar
		loginResult, adminLoggedOut := Login(sign, helo) // <-- Perubahan di sini: tangkap adminLoggedOut

		// Logika berdasarkan hasil login
		if loginResult {
			// Login berhasil, keluar dari loop
			running = false
		} else {
			// Login gagal, daftar selesai, atau admin logout - tampilkan pesan dan lanjutkan loop
			if adminLoggedOut { // <-- Perubahan di sini: tambahkan kondisi untuk adminLoggedOut
				fmt.Println("Anda telah berhasil logout.")
				// Tidak perlu menampilkan pesan "Login gagal" atau "Kembali ke menu login"
				// Langsung kembali ke loop utama untuk menampilkan menu login/daftar
			} else if *sign == "1" || *sign == "masuk" {
				fmt.Println("Login gagal. Coba lagi untuk masuk atau daftar.")
			} else if *sign == "2" || *sign == "daftar" {
				fmt.Println("Kembali ke menu login.")
			}
			atribut.Loading(200) // Beri jeda sebelum menampilkan menu lagi
		}
	}
}

// GetUsers mengembalikan array dan jumlah pengguna yang terdaftar (TIDAK MENGGUNAKAN SLICE)
func GetUsers() (Daftarnama, int) {
	return users, jumlahPengguna
}

// GetJumlahPengguna mengembalikan jumlah pengguna saat ini
func GetJumlahPengguna() int {
	return jumlahPengguna
}

// DeleteUser menghapus pengguna berdasarkan username
func DeleteUser(username string) bool {
	foundIndex := -1 // Inisialisasi dengan nilai yang tidak mungkin
	i := 0
	for i < jumlahPengguna {
		if users[i].Username == username {
			foundIndex = i
		}
		i++
	}

	if foundIndex != -1 { // Jika pengguna ditemukan
		// Geser semua elemen setelah foundIndex ke kiri
		j := foundIndex
		for j < jumlahPengguna-1 {
			users[j] = users[j+1]
			j++
		}
		// Kosongkan elemen terakhir (bekas duplikat)
		users[jumlahPengguna-1] = User{}
		jumlahPengguna--
		return true
	}
	return false
}

// GetUserByUsername mendapatkan pengguna berdasarkan username
func GetUserByUsername(username string) (User, bool) {
	var user User
	found := false
	i := 0
	for i < jumlahPengguna {
		if users[i].Username == username {
			user = users[i]
			found = true
		}
		i++
	}
	return user, found
}

// UpdateUser memperbarui informasi pengguna
func UpdateUser(oldUsername, newUsername, newPassword string) bool {
	updated := false
	i := 0
	for i < jumlahPengguna {
		if users[i].Username == oldUsername {
			users[i].Username = newUsername
			users[i].Password = newPassword
			updated = true
		}
		i++
	}
	return updated
}

// AddUser menambahkan pengguna baru ke dalam sistem
func AddUser(newUser User) bool {
	if jumlahPengguna < NMAX {
		// Cek duplikasi username sebelum menambah (sudah dilakukan di Login, tapi bagus juga di sini)
		if !isUsernameExists(newUser.Username) {
			users[jumlahPengguna] = newUser
			jumlahPengguna++
			return true
		}
	}
	return false
}

//==================================================================================
//                              ADMIN PERMISSIONS
//==================================================================================

const MAX_RIWAYAT int = 100

type RiwayatAdmin struct {
	Timestamp string
	Aktivitas string
	OlehAdmin string
}

var adminRiwayat [MAX_RIWAYAT]RiwayatAdmin
var jumlahRiwayat int

func MainAdmin(adminUsername string, logout *bool) {
	var pilihan string
	userFig := figure.NewFigure("ADMIN", "doom", true).String()
	welcomeFig := figure.NewFigure("WELCOME", "doom", true).String()

	running := true
	for running {
		atribut.ClearScreen()
		fmt.Print("\033[31m")
		fmt.Print(welcomeFig)
		fmt.Print(userFig)
		fmt.Print("\033[0m")
		fmt.Println(`
=========================================
          ADMIN PERMISSIONS
=========================================
1. Kelola Pengguna (Tambah/Hapus/Ubah)
2. Cetak Daftar Pengguna
3. Lihat Riwayat Aktivitas Admin
4. Reset Password Pengguna
5. Ganti Password Admin
6. Logout
0. Keluar Aplikasi
        `)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		if pilihan == "1" {
			KelolaPengguna(adminUsername)
		} else if pilihan == "2" {
			CetakPengguna()
		} else if pilihan == "3" {
			LihatRiwayatAdmin()
		} else if pilihan == "4" {
			ResetPasswordPengguna(adminUsername)
		} else if pilihan == "5" {
			GantiPasswordAdmin(adminUsername)
		} else if pilihan == "6" {
			fmt.Println("Anda telah logout dari mode admin.")
			running = false
			fmt.Println("\nTekan Enter ...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			*logout = true // Set flag logout ke true
		} else if pilihan == "0" {
			fmt.Println("Keluar dari aplikasi.")
			os.Exit(0)
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(200)
		}
	}
}

func CetakPengguna() {
	usersList, jumlah := GetUsers() // Menggunakan GetUsers yang mengembalikan array dan jumlah

	atribut.ClearScreen()
	fmt.Println("\n--- Daftar Pengguna ---")
	if jumlah == 0 {
		fmt.Println("Tidak ada pengguna terdaftar.")
	} else {
		i := 0
		for i < jumlah {
			fmt.Printf("%d. Username: %s\n", i+1, usersList[i].Username)
			i++
		}
	}
	fmt.Println("\nTekan Enter untuk kembali...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func KelolaPengguna(adminUsername string) {
	reader := bufio.NewReader(os.Stdin)

	managing := true
	for managing {
		atribut.ClearScreen()
		fmt.Println(`
=========================================
            KELOLA PENGGUNA
=========================================
1. Tambah Pengguna Baru
2. Hapus Pengguna
3. Ubah Data Pengguna
0. Kembali ke Menu Admin
        `)
		fmt.Print("Masukkan pilihan: ")
		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		if pilihan == "1" {
			TambahPenggunaBaru(adminUsername)
		} else if pilihan == "2" {
			HapusPengguna(adminUsername)
		} else if pilihan == "3" {
			UbahDataPengguna(adminUsername)
		} else if pilihan == "0" {
			managing = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(200)
		}
	}
}

func TambahPenggunaBaru(adminUsername string) bool {
	var berhasil bool = true
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Username baru: ")
	usernameInput, _ := reader.ReadString('\n')
	username := strings.TrimSpace(usernameInput)

	if isUsernameExists(username) { // Memanggil isUsernameExists
		fmt.Println("Username sudah ada. Harap gunakan username lain.")
		atribut.Loading(200)
		berhasil = false
	}

	fmt.Print("Masukkan Password untuk pengguna baru: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca kata sandi:", err)
		atribut.Loading(200)
		berhasil = false
	}
	password := string(passwordBytes)

	newUser := User{Username: username, Password: password}
	if AddUser(newUser) {
		fmt.Printf("Pengguna '%s' berhasil ditambahkan.\n", username)
		CatatRiwayatAdmin(adminUsername, fmt.Sprintf("Menambahkan pengguna baru: %s", username))
		berhasil = true
	} else {
		fmt.Println("Gagal menambahkan pengguna. Kapasitas penuh atau masalah lainnya.")
		berhasil = false
	}
	atribut.Loading(200)
	return berhasil
}

func HapusPengguna(adminUsername string) bool {
	var berhasil bool = true
	reader := bufio.NewReader(os.Stdin)
	CetakPengguna()
	fmt.Print("Masukkan username pengguna yang ingin dihapus: ")
	usernameInput, _ := reader.ReadString('\n')
	username := strings.TrimSpace(usernameInput)

	if username == "admin" && adminUsername == "admin" {
		fmt.Println("Anda tidak dapat menghapus akun admin utama!")
		atribut.Loading(200)
		berhasil = false
	}

	if DeleteUser(username) {
		fmt.Printf("Pengguna '%s' berhasil dihapus.\n", username)
		CatatRiwayatAdmin(adminUsername, fmt.Sprintf("Menghapus pengguna: %s", username))
		berhasil = true
	} else {
		fmt.Printf("Pengguna '%s' tidak ditemukan.\n", username)
		berhasil = false
	}
	atribut.Loading(200)
	return berhasil
}

func UbahDataPengguna(adminUsername string) bool {
	var berhasil bool = true
	reader := bufio.NewReader(os.Stdin)
	CetakPengguna()
	fmt.Print("Masukkan username pengguna yang ingin diubah: ")
	oldUsernameInput, _ := reader.ReadString('\n')
	oldUsername := strings.TrimSpace(oldUsernameInput)

	user, found := GetUserByUsername(oldUsername)
	if !found {
		fmt.Printf("Pengguna '%s' tidak ditemukan.\n", oldUsername)
		atribut.Loading(200)
		berhasil = false
	}

	fmt.Print("Masukkan Username baru (kosongkan jika tidak ingin mengubah): ")
	newUsernameInput, _ := reader.ReadString('\n')
	newUsername := strings.TrimSpace(newUsernameInput)
	if newUsername == "" {
		newUsername = user.Username
	} else if newUsername != oldUsername {
		if isUsernameExists(newUsername) { // Memanggil isUsernameExists
			fmt.Println("Username baru sudah digunakan. Harap pilih username lain.")
			atribut.Loading(200)
			berhasil = false
		}
	}

	fmt.Print("Masukkan Password baru (kosongkan jika tidak ingin mengubah): ")
	newPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca kata sandi:", err)
		atribut.Loading(200)
		berhasil = false
	}
	newPassword := string(newPasswordBytes)
	if newPassword == "" {
		newPassword = user.Password
	}

	if UpdateUser(oldUsername, newUsername, newPassword) {
		fmt.Printf("Data pengguna '%s' berhasil diubah menjadi Username: '%s', Password: '%s'\n", oldUsername, newUsername, newPassword)
		CatatRiwayatAdmin(adminUsername, fmt.Sprintf("Mengubah data pengguna dari '%s' menjadi '%s'", oldUsername, newUsername))
		berhasil = true
	} else {
		fmt.Println("Gagal mengubah data pengguna.")
		berhasil = false
	}
	atribut.Loading(200)
	return berhasil
}

func ResetPasswordPengguna(adminUsername string) bool {
	var berhasil bool = true
	reader := bufio.NewReader(os.Stdin)
	CetakPengguna()
	fmt.Print("Masukkan username pengguna yang ingin direset passwordnya: ")
	usernameInput, _ := reader.ReadString('\n')
	username := strings.TrimSpace(usernameInput)

	_, found := GetUserByUsername(username)
	if !found {
		fmt.Printf("Pengguna '%s' tidak ditemukan.\n", username)
		atribut.Loading(200)
		berhasil = false
	}

	fmt.Print("Masukkan password baru untuk pengguna ini: ")
	newPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca kata sandi:", err)
		atribut.Loading(200)
		berhasil = false
	}
	newPassword := string(newPasswordBytes)

	if UpdateUser(username, username, newPassword) {
		fmt.Printf("Password pengguna '%s' berhasil direset.\n", username)
		CatatRiwayatAdmin(adminUsername, fmt.Sprintf("Reset password pengguna: %s", username))
		berhasil = true
	} else {
		fmt.Println("Gagal mereset password pengguna.")
		berhasil = false
	}
	atribut.Loading(200)
	return berhasil
}

func GantiPasswordAdmin(adminUsername string) bool {
	var berhasil bool = true
	fmt.Print("Masukkan password admin lama Anda: ")
	oldPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca kata sandi lama:", err)
		atribut.Loading(200)
		berhasil = false
	}
	oldPassword := string(oldPasswordBytes)

	if !AuthenticateUser(adminUsername, oldPassword) {
		fmt.Println("Password lama Anda salah.")
		atribut.Loading(200)
		berhasil = false
	}

	fmt.Print("Masukkan password admin baru: ")
	newPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca kata sandi baru:", err)
		atribut.Loading(200)
		berhasil = false
	}
	newPassword := string(newPasswordBytes)

	fmt.Print("Konfirmasi password admin baru: ")
	confirmNewPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Println("Gagal membaca konfirmasi password baru:", err)
		atribut.Loading(200)
		berhasil = false
	}
	confirmNewPassword := string(confirmNewPasswordBytes)

	if newPassword != confirmNewPassword {
		fmt.Println("Password baru dan konfirmasi tidak cocok.")
		atribut.Loading(200)
		berhasil = false
	}

	if UpdateUser(adminUsername, adminUsername, newPassword) {
		fmt.Println("Password admin berhasil diubah.")
		CatatRiwayatAdmin(adminUsername, "Mengubah password akun admin sendiri")
		berhasil = true
	} else {
		fmt.Println("Gagal mengubah password admin.")
		berhasil = false
	}
	atribut.Loading(200)
	return berhasil
}

func CatatRiwayatAdmin(adminUsername, aktivitas string) {
	if jumlahRiwayat < MAX_RIWAYAT {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		riwayatEntry := RiwayatAdmin{
			Timestamp: timestamp,
			Aktivitas: aktivitas,
			OlehAdmin: adminUsername,
		}
		adminRiwayat[jumlahRiwayat] = riwayatEntry
		jumlahRiwayat++
	} else {
		fmt.Println("Riwayat admin penuh, tidak dapat mencatat aktivitas baru.")
	}
}

func LihatRiwayatAdmin() {
	atribut.ClearScreen()
	fmt.Println("\n--- Riwayat Aktivitas Admin ---")
	if jumlahRiwayat == 0 {
		fmt.Println("Tidak ada riwayat aktivitas admin.")
	} else {
		i := 0
		for i < jumlahRiwayat {
			entry := adminRiwayat[i]
			fmt.Printf("%d. [%s] Oleh: %s - %s\n", i+1, entry.Timestamp, entry.OlehAdmin, entry.Aktivitas)
			i++
		}
	}
	fmt.Println("\nTekan Enter untuk kembali...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
