package entity

import "fmt"

type BookedTicket struct {
	ID          int    `json:"id,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	OrderID     int    `json:"order_id,omitempty"`
	TrainID     int    `json:"train_id,omitempty"`
	WagonID     int    `json:"wagon_id,omitempty"`
	SeatIDs     []int  `json:"seat_ids,omitempty"`
	SeatIDsStr  string `json:"seat_ids_str,omitempty"`
	BookingDate string `json:"booking_date,omitempty"`
}

func (b *BookedTicket) Validate() error {
	if b.BookingDate == "" {
		return fmt.Errorf("booking date is empty")
	}

	if b.UserID == 0 {
		return fmt.Errorf("user id is empty")
	}

	if b.OrderID == 0 {
		return fmt.Errorf("order id is empty")
	}

	if b.TrainID == 0 {
		return fmt.Errorf("train id is empty")
	}

	if b.WagonID == 0 {
		return fmt.Errorf("wagon id is empty")
	}

	if len(b.SeatIDs) == 0 {
		return fmt.Errorf("seat id is empty")
	}

	b.SeatIDsStr = fmt.Sprintf("%v", b.SeatIDs)

	return nil
}
