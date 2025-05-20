package service

import (
	"bufio"
	"fmt"
	"latihan/data"
	"latihan/model"
	"latihan/utils"
	"os"
	"strings"
)

func OrderTicket() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Anda: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan Tujuan: ")
	destination, _ := reader.ReadString('\n')
	destination = strings.TrimSpace(destination)

	if name == "" || destination == "" {
		fmt.Println("Nama dan tujuan tidak boleh kosong!")
		return
	}

	tickets, err := data.LoadTickets()
	if err != nil {
		fmt.Println("Gagal memuat tiket:", err)
		return
	}

	for i, t := range tickets {
		if strings.EqualFold(t.Destination, destination) {
			if t.Stock <= 0 {
				fmt.Println("Maaf, tiket tujuan ini sudah habis.")
				return
			}

			tickets[i].Stock -= 1
			order := model.Order{
				ID:          utils.GenerateID(),
				Name:        name,
				Destination: t.Destination,
				Price:       t.Price,
			}

			orders, _ := data.LoadOrders()
			orders = append(orders, order)

			data.SaveTickets(tickets)
			data.SaveOrders(orders)

			fmt.Println("Selamat, pemesanan tiket berhasil!")
			return
		}
	}

	fmt.Println("Tujuan tidak ditemukan.")
}
