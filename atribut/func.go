package atribut

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Quiz struct {
	Nama                         string
	ID                           string
	TotalScore                   int
	GoLanguageScore              int
	PercabanganScore             int
	KonversiTipeDataScore        int
	OperasiMatematikaLogikaScore int
	PerulanganScore              int
	TipeDataGoScore              int
	VariableConstantScore        int
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Peringatan: Gagal membersihkan layar. Error: %v\n", err)
	}
}
func Loading(ms int) {
	fmt.Print("Loading...")
	time.Sleep(time.Duration(ms) * time.Millisecond)
	fmt.Println()
}

// GetWaktuSekarang mengembalikan string representasi waktu saat ini (YYYY-MM-DD HH:MM:SS)
func GetWaktuSekarang() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
