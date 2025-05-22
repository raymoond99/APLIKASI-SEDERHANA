package sampah

import (
	"fmt"
	"strings"
)

func Sortir() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk disortir.")
		return
	}

	fmt.Println("\nUrutkan berdasarkan:")
	fmt.Println("1. Nama (A-Z) [Insertion Sort]")
	fmt.Println("2. Nama (Z-A) [Insertion Sort]")
	fmt.Println("3. Jumlah (Terkecil) [Selection Sort]")
	fmt.Println("4. Jumlah (Terbesar) [Selection Sort]")

	pilihan := Input("Pilih opsi sortir (1-4): ")

	switch pilihan {
	case "1":
		insertionSortNama(true)
	case "2":
		insertionSortNama(false)
	case "3":
		selectionSortJumlah(true)
	case "4":
		selectionSortJumlah(false)
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("Data berhasil disortir.")
	Lihat()
}

func insertionSortNama(asc bool) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && compareNama(data[j].Nama, key.Nama, asc) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func compareNama(a, b string, asc bool) bool {
	if asc {
		return strings.ToLower(a) > strings.ToLower(b)
	}
	return strings.ToLower(a) < strings.ToLower(b)
}

func selectionSortJumlah(asc bool) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if compareJumlah(data[j].Jumlah, data[idx].Jumlah, asc) {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

func compareJumlah(a, b float64, asc bool) bool {
	if asc {
		return a < b
	}
	return a > b
}
