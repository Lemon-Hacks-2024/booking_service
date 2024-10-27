package entity

type Wagon struct {
	WagonID        int    `json:"wagon_id"`
	Type           string `json:"type"`
	AvailableSeats int    `json:"available_seats"`
	Seats          []Seat `json:"seats"`
}

func (w *Wagon) UpdateAvailableSeats() {
	for _, seat := range w.Seats {
		if seat.BookingStatus == "FREE" {
			w.AvailableSeats++
		}
	}
}

type WagonResponse struct {
	WagonID int    `json:"wagon_id"`
	Type    string `json:"type"`
	Seats   []struct {
		SeatID        int    `json:"seat_id"`
		SeatNum       string `json:"seatNum"`
		Block         string `json:"block"`
		Price         int    `json:"price"`
		BookingStatus string `json:"bookingStatus"`
	} `json:"seats"`
}
