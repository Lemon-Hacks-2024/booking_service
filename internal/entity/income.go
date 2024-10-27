package entity

type Income struct {
	UserID        int      `json:"user_id"`
	TrainID       int      `json:"train_id,omitempty"`
	WagonID       int      `json:"wagon_id,omitempty"`
	SeatID        int      `json:"seat_id,omitempty"`
	Route         string   `json:"route"`
	DateFrom      string   `json:"date_from"`
	DateTo        string   `json:"date_to"`
	WagonType     string   `json:"wagon_type,omitempty"`
	PlacePosition []string `json:"place_position,omitempty"`
	Price         int      `json:"price,omitempty"`
	SeatsQty      int      `json:"seats_qty,omitempty"`
	NeedNearby    bool     `json:"need_nearby,omitempty"`
}
