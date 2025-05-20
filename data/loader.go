package data

import (
	"encoding/json"
	"latihan/model"
	"os"
)

func LoadTickets() ([]model.Ticket, error) {
	file, err := os.ReadFile("tiket.json")
	if err != nil {
		return nil, err
	}
	var tickets []model.Ticket
	err = json.Unmarshal(file, &tickets)
	return tickets, err
}

func SaveTickets(tickets []model.Ticket) error {
	data, err := json.MarshalIndent(tickets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tiket.json", data, 0644)
}

func LoadOrders() ([]model.Order, error) {
	file, err := os.ReadFile("order.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Order{}, nil
		}
		return nil, err
	}
	var orders []model.Order
	err = json.Unmarshal(file, &orders)
	return orders, err
}

func SaveOrders(orders []model.Order) error {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("order.json", data, 0644)
}
