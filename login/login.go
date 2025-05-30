package login

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type user struct {
	Username, Password string
}

const nmax int = 43

type daftarNama [nmax]user

var users daftarNama
var jumlahPengguna int = 42

func Login() {
	var username, password, sign string
	var signup user
	stop := true
	key := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%15s\n", "==== SIGNIN ====")
	fmt.Printf("%15s\n", "==== SIGNUP ====")
	fmt.Print("pilih yang mana: ")
	signed, _ := reader.ReadString('\n')
	sign = strings.TrimSpace(signed)

	if sign == "signin" {
		for stop && key < 3 {
			fmt.Print("Masukkan Username: ")
			usernameInput, _ := reader.ReadString('\n')
			username = strings.TrimSpace(usernameInput)

			fmt.Print("Masukkan Password: ")
			passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if err != nil {
				fmt.Println("Gagal membaca kata sandi:", err)
				return
			}
			password = string(passwordBytes)

			if authenticateUser(username, password) {
				fmt.Println("Login berhasil!")
				stop = false
			} else {
				fmt.Println("Username atau password salah.")
				key++
				fmt.Printf("Tersisa %d kesempatan\n", 3-key)
			}
		}
	} else if sign == "signup" {
		if jumlahPengguna >= nmax {
			fmt.Println("Maaf, jumlah maksimum pengguna telah tercapai.")
			return
		}

		fmt.Print("Masukkan Username: ")
		signup.Username, _ = reader.ReadString('\n')
		signup.Username = strings.TrimSpace(signup.Username)

		fmt.Print("Masukkan NIM: ")
		NIM, _ := reader.ReadString('\n')
		NIM = strings.TrimSpace(NIM)

		fmt.Println("Selamat datang,")
		fmt.Println("Nama:", signup.Username)
		fmt.Println("NIM:", NIM)

		fmt.Print("Masukkan Password: ")
		passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			fmt.Println("Gagal membaca kata sandi:", err)
			return
		}

		fmt.Print("Konfirmasi Password: ")
		confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			fmt.Println("Gagal membaca konfirmasi password:", err)
			return
		}

		if string(passwordBytes) != string(confirmBytes) {
			fmt.Println("Password dan konfirmasi tidak cocok.")
			return
		}

		signup.Password = string(passwordBytes)
		users[jumlahPengguna] = signup
		jumlahPengguna++
		fmt.Println("Signup berhasil! Silakan login.")
	} else {
		fmt.Println("Pilihan tidak dikenali.")
	}
}

func authenticateUser(username, password string) bool {
	for i := 0; i < jumlahPengguna; i++ {
		if users[i].Username == username && users[i].Password == password {
			return true
		}
	}
	return false
}

func initUsers() {
	users[0] = user{"admin", "admin123"}
	users[1] = user{"nathasyayuanmaharani", "0001"}
	users[2] = user{"theodoreelvisestrada", "0002"}
	users[3] = user{"dyahkusumawardani", "0003"}
	users[4] = user{"azrielraihaneldovahartoto", "0009"}
	users[5] = user{"muhammadilhamalifianda", "0010"}
	users[6] = user{"alyaazizaputeri", "0012"}
	users[7] = user{"ahmadabdansyakuro", "0013"}
	users[8] = user{"fathurrahmanalfarizi", "0014"}
	users[9] = user{"nuswantorosetyomukti", "0015"}
	users[10] = user{"anggitacahyatihidayat", "0016"}
	users[11] = user{"wibnuhijrahfranstio", "0017"}
	users[12] = user{"meyshaprimiandita", "0018"}
	users[13] = user{"muhamadfiqrihabibi", "0032"}
	users[14] = user{"fitriacahyani", "0034"}
	users[15] = user{"triansyahdaniswaraibrahim", "0066"}
	users[16] = user{"rakhaabdillahalkautsar", "0068"}
	users[17] = user{"avicenanaufallathif", "0071"}
	users[18] = user{"naylaassyifa", "0072"}
	users[19] = user{"williampetervanxnajoan", "0084"}
	users[20] = user{"rayvanalifarlomahesworo", "0087"}
	users[21] = user{"zaidansalamrojab", "0093"}
	users[22] = user{"audreyfredileyhanas", "0099"}
	users[23] = user{"muhammadnaelfadly", "0096"}
	users[24] = user{"nairacahayaputridarmawansinaga", "0100"}
	users[25] = user{"muhamadalwansuryadi", "0104"}
	users[26] = user{"dhafyahmadzubaidi", "0117"}
	users[27] = user{"muhammadfarisdhiyaylhaqsarbini", "0119"}
	users[28] = user{"nursyadira", "0123"}
	users[29] = user{"rayfitokrisnawijaya", "0124"}
	users[30] = user{"mochammadrafirisqullah", "0125"}
	users[31] = user{"iputugedeagastyakrisnawidartha", "0126"}
	users[32] = user{"rendil", "0137"}
	users[33] = user{"muhammadariqazzaki", "0138"}
	users[34] = user{"edmundyuliusgantur", "0152"}
	users[35] = user{"muhammadsayyidhuwaidi", "0153"}
	users[36] = user{"muhdzuljalalwaliikramjalil", "0157"}
	users[37] = user{"ramadhantangguhdefennder", "0301"}
	users[38] = user{"adzkiyaputrirahmawan", "0329"}
	users[39] = user{"fathimahradhiyya", "0332"}
	users[40] = user{"rakanghazianadiwjaya", "0336"}
	users[41] = user{"jihannabilamubarakah", "0037"}
	// Jadi kamu perlu ubah totalUser jadi 42 kalau mau tambah ini
}

func Mainlogin() {
	var stop bool = true
	for stop {
		initUsers()
		Login()
	}
}
