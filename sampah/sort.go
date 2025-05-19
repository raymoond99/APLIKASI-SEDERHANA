package sampah

import (
	"fmt"
	"sort"
	"strings"
)

func Sortir() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk disortir.")
		return
	}

	fmt.Println("\nUrutkan berdasarkan:")
	fmt.Println("1. Nama (A-Z)")
	fmt.Println("2. Nama (Z-A)")
	fmt.Println("3. Jumlah (Terkecil)")
	fmt.Println("4. Jumlah (Terbesar)")

	pilihan := Input("Pilih opsi sortir (1-4): ")

	switch pilihan {
	case "1":
		sort.Slice(data, func(i, j int) bool {
			return strings.ToLower(data[i].Nama) < strings.ToLower(data[j].Nama)
		})
	case "2":
		sort.Slice(data, func(i, j int) bool {
			return strings.ToLower(data[i].Nama) > strings.ToLower(data[j].Nama)
		})
	case "3":
		sort.Slice(data, func(i, j int) bool {
			return data[i].Jumlah < data[j].Jumlah
		})
	case "4":
		sort.Slice(data, func(i, j int) bool {
			return data[i].Jumlah > data[j].Jumlah
		})
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("Data berhasil disortir.")
	Lihat()
}
