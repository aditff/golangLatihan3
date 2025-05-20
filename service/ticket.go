package service

import (
	"fmt"
	"latihan/data"
)

func ShowTickets() {
	tickets, err := data.LoadTickets()
	if err != nil {
		fmt.Println("Gagal memuat data tiket:", err)
		return
	}

	fmt.Println("\n=== Daftar Tiket Tersedia ===")
	for _, t := range tickets {
		fmt.Printf("Tujuan: %-10s | Harga: Rp%.0f | Stok: %d\n", t.Destination, t.Price, t.Stock)
	}
	fmt.Println("==============================")
}
