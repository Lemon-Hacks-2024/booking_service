package service

import (
	"booking_service/internal/entity"
	"booking_service/internal/repository"
)

type BookedTicketService struct {
	bookedTicketRepository repository.BookedTicket
}

func NewBookedTicketService(bookedTicketRepository *repository.Repository) *BookedTicketService {
	return &BookedTicketService{
		bookedTicketRepository: bookedTicketRepository,
	}
}

func (s *BookedTicketService) Create(ticket entity.BookedTicket) (int, error) {
	return s.bookedTicketRepository.Create(ticket)
}
