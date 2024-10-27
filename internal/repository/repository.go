package repository

import (
	"booking_service/internal/entity"
	"booking_service/pkg"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	GetUserByID(userID int) (entity.User, error)
	GetUserByLogIN(email, password string) (entity.User, error)
}

type BookedTicket interface {
	Create(ticket entity.BookedTicket) (int, error)
}

type Repository struct {
	User
	BookedTicket
}

func NewRepository(postgres *sqlx.DB, redis *pkg.Redis) *Repository {
	return &Repository{
		User:         NewUserRepository(postgres),
		BookedTicket: NewBookedTicketRepository(postgres),
	}
}
