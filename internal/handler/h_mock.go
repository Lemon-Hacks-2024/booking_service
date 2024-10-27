package handler

import "github.com/gin-gonic/gin"

type ModelMockTrains []struct {
	TrainID             int    `json:"train_id"`
	GlobalRoute         string `json:"global_route"`
	StartpointDeparture string `json:"startpoint_departure"`
	EndpointArrival     string `json:"endpoint_arrival"`
	DetailedRoute       []struct {
		Name      string      `json:"name"`
		Num       int         `json:"num"`
		Arrival   interface{} `json:"arrival"`
		Departure interface{} `json:"departure"`
	} `json:"detailed_route"`
	WagonsInfo []struct {
		WagonID int    `json:"wagon_id"`
		Type    string `json:"type"`
	} `json:"wagons_info"`
	AvailableSeatsCount int `json:"available_seats_count"`
}

func (h *Handler) MockTrains(ctx *gin.Context) {

	ctx.JSON(200, ModelMockTrains{
		{
			TrainID:             1,
			GlobalRoute:         "Ростов-на-Дону->Москва",
			StartpointDeparture: "25.10.2024 11:07:00",
			EndpointArrival:     "26.10.2024 13:09:00",
			DetailedRoute: []struct {
				Name      string      `json:"name"`
				Num       int         `json:"num"`
				Arrival   interface{} `json:"arrival"`
				Departure interface{} `json:"departure"`
			}{
				{
					Name:      "Ростов-на-Дону",
					Num:       1,
					Arrival:   nil,
					Departure: "25.10.2024 11:07:00",
				},
				{
					Name:      "Москва",
					Num:       1,
					Arrival:   "26.10.2024 13:09:00",
					Departure: nil,
				},
			},
			WagonsInfo: []struct {
				WagonID int    `json:"wagon_id"`
				Type    string `json:"type"`
			}{
				{
					WagonID: 1,
					Type:    "PLATZCART",
				},
				{
					WagonID: 2,
					Type:    "COUPE",
				},
			},
			AvailableSeatsCount: 12,
		},
		{
			TrainID:             3,
			GlobalRoute:         "Ростов-на-Дону->Москва",
			StartpointDeparture: "28.10.2024 11:07:00",
			EndpointArrival:     "29.10.2024 13:09:00",
			DetailedRoute: []struct {
				Name      string      `json:"name"`
				Num       int         `json:"num"`
				Arrival   interface{} `json:"arrival"`
				Departure interface{} `json:"departure"`
			}{
				{
					Name:      "Ростов-на-Дону",
					Num:       1,
					Arrival:   nil,
					Departure: "28.10.2024 11:07:00",
				},
				{
					Name:      "Москва",
					Num:       2,
					Arrival:   "29.10.2024 13:09:00",
					Departure: nil,
				},
			},
			WagonsInfo: []struct {
				WagonID int    `json:"wagon_id"`
				Type    string `json:"type"`
			}{
				{
					WagonID: 5,
					Type:    "PLATZCART",
				},
				{
					WagonID: 6,
					Type:    "COUPE",
				},
			},
			AvailableSeatsCount: 1,
		},
		{
			TrainID:             4,
			GlobalRoute:         "Ростов-на-Дону->Москва",
			StartpointDeparture: "25.10.2024 11:07:00",
			EndpointArrival:     "26.10.2024 13:09:00",
			DetailedRoute: []struct {
				Name      string      `json:"name"`
				Num       int         `json:"num"`
				Arrival   interface{} `json:"arrival"`
				Departure interface{} `json:"departure"`
			}{
				{
					Name:      "Ростов-на-Дону",
					Num:       1,
					Arrival:   nil,
					Departure: "25.10.2024 11:07:00",
				},
				{
					Name:      "Москва",
					Num:       2,
					Arrival:   "26.10.2024 13:09:00",
					Departure: nil,
				},
			},
			WagonsInfo: []struct {
				WagonID int    `json:"wagon_id"`
				Type    string `json:"type"`
			}{
				{
					WagonID: 3,
					Type:    "PLATZCART",
				},
				{
					WagonID: 4,
					Type:    "COUPE",
				},
			},
			AvailableSeatsCount: 3,
		},
	})
	return
}

type ModelMockSeats []struct {
	SeatID        int    `json:"seat_id"`
	SeatNum       string `json:"seatNum"`
	Block         string `json:"block"`
	Price         int    `json:"price"`
	BookingStatus string `json:"bookingStatus"`
}

func (h *Handler) MockSeats(ctx *gin.Context) {
	ctx.JSON(200, ModelMockSeats{
		{
			SeatID:        1,
			SeatNum:       "1",
			Block:         "1",
			Price:         300,
			BookingStatus: "FREE",
		},
		{
			SeatID:        2,
			SeatNum:       "2",
			Block:         "1",
			Price:         300,
			BookingStatus: "FREE",
		},
		{
			SeatID:        3,
			SeatNum:       "3",
			Block:         "1",
			Price:         300,
			BookingStatus: "FREE",
		},
		{
			SeatID:        4,
			SeatNum:       "4",
			Block:         "1",
			Price:         300,
			BookingStatus: "FREE",
		},
		{
			SeatID:        5,
			SeatNum:       "5",
			Block:         "1",
			Price:         300,
			BookingStatus: "FREE",
		},
	})
	return
}
