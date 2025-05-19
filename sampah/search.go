package sampah

import (
	"fmt"
	"strings"
)

func Cari() {
	if len(data) == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	kunci := strings.ToLower(Input("Masukkan kata kunci nama sampah: "))
	ditemukan := false

	for i, d := range data {
		if strings.Contains(strings.ToLower(d.Nama), kunci) {
			if !ditemukan {
				fmt.Println("Hasil pencarian:")
			}
			fmt.Printf("%d. %s - %.2f kg\n", i+1, d.Nama, d.Jumlah)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan data yang cocok.")
	}
}
