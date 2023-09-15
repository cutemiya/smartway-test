package ticket

import (
	"smartway-test/client/doc"
	"smartway-test/client/user"
)

type Ticket struct {
	StartPoint string `json:"startPoint"`
	EndPoint   string `json:"endPoint"`

	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`

	Company string `json:"company"`
}

type FullInfoTicket struct {
	Ticket
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	BuyTime string `json:"buyTime"`
}

type FullInfoTicketList struct {
	Tickets []FullInfoTicket `json:"tickets"`
}

type PassengersByTicketId struct {
	Passengers []user.User `json:"passengers"`
}

type Passenger struct {
	user.User
	Documents []doc.Document `json:"documents"`
}

type AllINfoAboutTicket struct {
	Ticket
	BuyTime    string      `json:"buyTime"`
	Passengers []Passenger `json:"passengers"`
}
