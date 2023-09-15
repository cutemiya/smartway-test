package model

type Ticket struct {
	StartPoint string
	EndPoint   string

	StartTime string
	EndTime   string

	Company string
}

type FullTicketInfo struct {
	Ticket
	Id      int
	UserId  int
	BuyTime string
}

type AllInfoOfTicket struct {
	Ticket
	BuyTime    string
	Passengers []Passenger
}
