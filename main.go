package main

import (
	"aplikasi-sederhana/sampah"
	"fmt"
)

func main() {
	sampah.Baca()
	defer sampah.Simpan()

	for {
		fmt.Println("\n====== MENU PENGELOLAAN SAMPAH ======")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Lihat Data")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Urutkan Data")
		fmt.Println("6. Cari Data")
		fmt.Println("7. Keluar")

		pilihan := sampah.Input("Pilih menu (1-7): ")

		switch pilihan {
		case "1":
			sampah.Tambah()
		case "2":
			sampah.Lihat()
		case "3":
			sampah.Edit()
		case "4":
			sampah.Hapus()
		case "5":
			sampah.Sortir()
		case "6":
			sampah.Cari()
		case "7":
			fmt.Println("Terima kasih! Data disimpan.")
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}
