// package dto

// type TicketRequest struct {
// 	Passenger   string
// 	Destination string
// }

// func NewRequest(passenger, destination string) TicketRequest {
// 	return TicketRequest{
// 		Passenger:   passenger,
// 		Destination: destination,
// 	}
// }

package dto

type TicketRequest struct {
	Name        string
	Destination string
}
