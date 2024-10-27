package service

import (
	"booking_service/internal/entity"
	"booking_service/internal/repository"
	"booking_service/pkg"
)

type User interface {
	CreateUser(user entity.User) (pkg.JWT, error)
	Login(email, password string) (pkg.JWT, error)
	GetUserByToken(token string) (entity.User, error)
}

type Train interface {
	GetAllTrains(queryParam entity.TrainsInputQueryParam) ([]entity.Train, error)
	GetWagonsByTrain(trainID int) ([]entity.Wagon, error)
}

type BookedTicket interface {
	Create(ticket entity.BookedTicket) (int, error)
}

type Service struct {
	User
	Train
	BookedTicket
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:         NewUserService(repos),
		Train:        NewAxTrainService(),
		BookedTicket: NewBookedTicketService(repos),
	}
}
