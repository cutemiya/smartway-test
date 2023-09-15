package ticket

import (
	"smartway-test/client/doc"
	"smartway-test/client/user"
	"smartway-test/model"
)

func MapToServiceTicketModel(ticket Ticket) model.Ticket {
	return model.Ticket{
		StartPoint: ticket.StartPoint,
		EndPoint:   ticket.EndPoint,
		StartTime:  ticket.StartTime,
		EndTime:    ticket.EndTime,
		Company:    ticket.Company,
	}
}

func MapToServiceFullTicket(ticket FullInfoTicket) model.FullTicketInfo {
	return model.FullTicketInfo{
		Id:      ticket.Id,
		UserId:  ticket.UserId,
		BuyTime: ticket.BuyTime,
		Ticket: model.Ticket{
			StartPoint: ticket.StartPoint,
			EndPoint:   ticket.EndPoint,
			StartTime:  ticket.StartTime,
			EndTime:    ticket.EndTime,
			Company:    ticket.Company,
		},
	}
}

// reverse map

func MapToClientTicketList(tickets []model.FullTicketInfo) FullInfoTicketList {
	var ticketList []FullInfoTicket

	for _, ticket := range tickets {
		ticketList = append(ticketList, FullInfoTicket{
			Id:      ticket.Id,
			UserId:  ticket.UserId,
			BuyTime: ticket.BuyTime,
			Ticket: Ticket{
				StartPoint: ticket.StartPoint,
				EndPoint:   ticket.EndPoint,
				StartTime:  ticket.StartTime,
				EndTime:    ticket.EndTime,
				Company:    ticket.Company,
			},
		})
	}

	return FullInfoTicketList{ticketList}
}

func MapToClientPassengers(servicePassengers []model.User) PassengersByTicketId {
	var passengers []user.User

	for _, passenger := range servicePassengers {
		passengers = append(passengers, user.User{
			Name:       passenger.Name,
			Surname:    passenger.Surname,
			Patronymic: passenger.Patronymic,
		})
	}

	return PassengersByTicketId{passengers}
}

func MapToClientAllInfoAboutTicket(allInfoModel model.AllInfoOfTicket) AllINfoAboutTicket {
	var passengers []Passenger

	for _, passenger := range allInfoModel.Passengers {
		var documents []doc.Document

		for _, d := range passenger.Documents {
			documents = append(documents, doc.Document{
				Id:     d.Id,
				Type:   d.Type,
				Number: d.Number,
			})
		}

		passengers = append(passengers, Passenger{
			User: user.User{
				Name:       passenger.Name,
				Surname:    passenger.Surname,
				Patronymic: passenger.Patronymic,
			},
			Documents: documents,
		})
	}

	return AllINfoAboutTicket{
		BuyTime: allInfoModel.BuyTime,
		Ticket: Ticket{
			StartPoint: allInfoModel.StartPoint,
			EndPoint:   allInfoModel.EndPoint,
			StartTime:  allInfoModel.StartTime,
			EndTime:    allInfoModel.EndTime,
			Company:    allInfoModel.Company,
		},
		Passengers: passengers,
	}
}
