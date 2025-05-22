package sampah

import (
	"fmt"
	"sort"
	"strings"
)

var data []Sampah

type Sampah struct {
	Nama   string
	Jumlah float64
}

func Input(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func Cari() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	metode := strings.ToLower(Input("Pilih metode pencarian (sequential/binary): "))
	kunci := strings.ToLower(Input("Masukkan kata kunci nama sampah: "))

	switch metode {
	case "sequential":
		sequentialSearch(kunci)
	case "binary":

		sort.Slice(data, func(i, j int) bool {
			return strings.ToLower(data[i].Nama) < strings.ToLower(data[j].Nama)
		})
		binarySearch(kunci)
	default:
		fmt.Println("Metode pencarian tidak dikenal.")
	}
}

func sequentialSearch(kunci string) {
	ditemukan := false
	for i, d := range data {
		if strings.Contains(strings.ToLower(d.Nama), kunci) {
			if !ditemukan {
				fmt.Println("Hasil pencarian (Sequential Search):")
			}
			fmt.Printf("%d. %s - %.2f kg\n", i+1, d.Nama, d.Jumlah)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ditemukan data yang cocok.")
	}
}

func binarySearch(kunci string) {
	low, high := 0, len(data)-1
	ditemukan := false
	fmt.Println("Hasil pencarian (Binary Search):")

	for low <= high {
		mid := (low + high) / 2
		namaLower := strings.ToLower(data[mid].Nama)

		if strings.Contains(namaLower, kunci) {
			fmt.Printf("%d. %s - %.2f kg\n", mid+1, data[mid].Nama, data[mid].Jumlah)
			ditemukan = true
			break
		}

		if kunci < namaLower {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan data yang cocok.")
	}
}
