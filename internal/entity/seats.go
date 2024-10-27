package entity

type Seat struct {
	SeatID        int    `json:"seat_id"`
	SeatNum       int    `json:"seat_num"`
	Block         int    `json:"block"`
	Price         int    `json:"price"`
	BookingStatus string `json:"booking_status"`
}
