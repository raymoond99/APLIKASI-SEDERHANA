package sampah

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type DataSampah struct {
	Nama   string
	Jumlah float64
}

var data []DataSampah

const namaFile = "data_sampah.txt"

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Baca() {
	file, err := os.ReadFile(namaFile)
	if err == nil {
		json.Unmarshal(file, &data)
	}
}

func Simpan() {
	file, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(namaFile, file, 0644)
}

func Tambah() {
	nama := Input("Masukkan nama sampah (misal: botol plastik): ")
	jumlahStr := Input("Jumlah (kg): ")

	var jumlah float64
	fmt.Sscanf(jumlahStr, "%f", &jumlah)

	data = append(data, DataSampah{Nama: nama, Jumlah: jumlah})

	fmt.Printf("Sampah dikenali sebagai: %s\n", DeteksiJenis(nama))
	fmt.Printf("Bisa didaur ulang dengan metode: %s\n", MetodeDaurUlang(nama))
	fmt.Println("Data ditambahkan!")
}

func Lihat() {
	if len(data) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, d := range data {
		fmt.Printf("%d. %s - %.2f kg\n", i+1, d.Nama, d.Jumlah)
	}
}

func Edit() {
	Lihat()
	iStr := Input("Masukkan nomor data yang ingin diedit: ")
	var i int
	fmt.Sscanf(iStr, "%d", &i)
	i--

	if i < 0 || i >= len(data) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	nama := Input("Nama baru: ")
	jumlahStr := Input("Jumlah baru (kg): ")
	var jumlah float64
	fmt.Sscanf(jumlahStr, "%f", &jumlah)

	data[i] = DataSampah{Nama: nama, Jumlah: jumlah}
	fmt.Println("Data berhasil diedit!")
}

func Hapus() {
	Lihat()
	iStr := Input("Masukkan nomor data yang ingin dihapus: ")
	var i int
	fmt.Sscanf(iStr, "%d", &i)
	i--

	if i < 0 || i >= len(data) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	data = append(data[:i], data[i+1:]...)
	fmt.Println("Data berhasil dihapus.")
}

func DeteksiJenis(nama string) string {
	n := strings.ToLower(nama)
	switch {
	case strings.Contains(n, "plastik"):
		return "Plastik"
	case strings.Contains(n, "kaca"):
		return "Kaca"
	case strings.Contains(n, "kertas"):
		return "Kertas"
	default:
		return "Lainnya"
	}
}

func MetodeDaurUlang(nama string) string {
	switch DeteksiJenis(nama) {
	case "Plastik":
		return "Dicuci dan dilelehkan"
	case "Kaca":
		return "Dilebur dan dicetak ulang"
	case "Kertas":
		return "Direndam dan dicetak kembali"
	default:
		return "Tidak tersedia"
	}
}
