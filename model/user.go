package model

type User struct {
	Name       string
	Surname    string
	Patronymic string
}

type Passenger struct {
	User
	Documents []Document
}

type ReportFlights struct {
	Previously                []FullTicketInfo
	NotFulFilled              []FullTicketInfo
	PreviouslyAndNotFulFilled []FullTicketInfo
}
