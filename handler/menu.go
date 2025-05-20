package handler

import (
	"bufio"
	"fmt"
	"latihan/service"
	"os"
)

func RunMenu() {
	for {
		fmt.Println("=== Sistem Tiket Online PT Gobus ===")
		fmt.Println("1. Lihat Daftar Tiket")
		fmt.Println("2. Pesan Tiket")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih Menu (1-3): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			service.ShowTickets()
		case 2:
			service.OrderTicket()
		case 3:
			fmt.Println("Terima kasih telah menggunakan sistem.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Println("\nTekan ENTER untuk kembali ke menu...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
