package repository

import (
	"booking_service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookedTicketRepository struct {
	postgres *sqlx.DB
}

func NewBookedTicketRepository(postgres *sqlx.DB) *BookedTicketRepository {
	return &BookedTicketRepository{
		postgres: postgres,
	}
}

func (r *BookedTicketRepository) Create(ticket entity.BookedTicket) (int, error) {
	var id int

	query := `INSERT INTO booked_tickets (user_id, order_id, train_id, wagon_id, seat_ids, booking_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	row := r.postgres.QueryRow(query, ticket.UserID, ticket.OrderID, ticket.TrainID, ticket.WagonID, ticket.SeatIDsStr, ticket.BookingDate)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
