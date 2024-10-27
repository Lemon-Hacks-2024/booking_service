package entity

import (
	"strconv"
	"time"
)

type Train struct {
	TrainID             int    `json:"train_id,omitempty"`
	Route               string `json:"route,omitempty"`
	StartPoint          string `json:"start_point,omitempty"`
	StartPointDeparture string `json:"start_point_departure,omitempty"`
	StartPointTimeMs    int64  `json:"start_point_time_ms,omitempty"`
	StartPointTime      string `json:"start_point_time,omitempty"`
	EndPoint            string `json:"endpoint,omitempty"`
	EndPointArrival     string `json:"endpoint_arrival,omitempty"`
	EndPointTimeMs      int64  `json:"endpoint_time_ms,omitempty"`
	EndPointTime        string `json:"endpoint_time,omitempty"`
	TravelTime          string `json:"travel_time,omitempty"`
	DetailedRoute       []struct {
		Name      string `json:"name,omitempty"`
		Num       int    `json:"num,omitempty"`
		Arrival   string `json:"arrival,omitempty"`
		Departure string `json:"departure,omitempty"`
	} `json:"detailed_route,omitempty"`
	WagonsInfo          []TrainsWagonsInfo `json:"wagons_info,omitempty"`
	AvailableSeatsCount int                `json:"available_seats_count"`
	Error               string             `json:"error,omitempty"`
}

func (t *Train) FillFields() {
	parsedDateDeparture, _ := time.Parse("02.01.2006 15:04:05", t.StartPointDeparture)
	t.StartPointTimeMs = parsedDateDeparture.Unix() * 1000

	parsedDateArrival, _ := time.Parse("02.01.2006 15:04:05", t.EndPointArrival)
	t.EndPointTimeMs = parsedDateArrival.Unix() * 1000

	startPointTime := ""
	if parsedDateDeparture.Hour() <= 9 {
		startPointTime += "0" + strconv.Itoa(parsedDateDeparture.Hour())
	} else {
		startPointTime += strconv.Itoa(parsedDateDeparture.Hour())
	}
	//
	if parsedDateDeparture.Minute() <= 9 {
		startPointTime += ":0" + strconv.Itoa(parsedDateDeparture.Minute())
	} else {
		startPointTime += ":" + strconv.Itoa(parsedDateDeparture.Minute())
	}
	//
	t.StartPointTime = startPointTime

	endPointTime := ""
	if parsedDateArrival.Hour() <= 9 {
		endPointTime += "0" + strconv.Itoa(parsedDateArrival.Hour())
	} else {
		endPointTime += strconv.Itoa(parsedDateArrival.Hour())
	}
	//
	if parsedDateArrival.Minute() <= 9 {
		endPointTime += ":0" + strconv.Itoa(parsedDateArrival.Minute())
	} else {
		endPointTime += ":" + strconv.Itoa(parsedDateArrival.Minute())
	}
	//
	t.EndPointTime = endPointTime

	travelTime := parsedDateArrival.Sub(parsedDateDeparture)
	if travelTime.Minutes() <= 9 {
		t.TravelTime = strconv.Itoa(int(travelTime.Hours())) + " ч 0" + strconv.Itoa(int(travelTime.Minutes())-int(travelTime.Hours())*60) + " мин"
	} else {
		t.TravelTime = strconv.Itoa(int(travelTime.Hours())) + " ч " + strconv.Itoa(int(travelTime.Minutes())-int(travelTime.Hours())*60) + " мин"
	}

}

type TrainsWagonsInfo struct {
	WagonID  int    `json:"wagon_id,omitempty"`
	Type     string `json:"type,omitempty"`
	Seats    int    `json:"seats,omitempty"`
	PriceMin int    `json:"price_min,omitempty"`
	PriceMax int    `json:"price_max,omitempty"`
}

type TrainsResponse []struct {
	TrainID             int    `json:"train_id"`
	GlobalRoute         string `json:"global_route"`
	StartPointDeparture string `json:"startpoint_departure"`
	EndpointArrival     string `json:"endpoint_arrival"`
	DetailedRoute       []struct {
		Name      string `json:"name"`
		Num       int    `json:"num"`
		Arrival   string `json:"arrival"`
		Departure string `json:"departure"`
	} `json:"detailed_route"`
	WagonsInfo []struct {
		WagonID int    `json:"wagon_id"`
		Type    string `json:"type"`
	} `json:"wagons_info"`
	AvailableSeatsCount int    `json:"available_seats_count"`
	Error               string `json:"error"`
}

type TrainsInputQueryParam struct {
	BookingAvailable   string    `json:"booking_available,omitempty"`
	StartPoint         string    `json:"start_point,omitempty"`
	EndPoint           string    `json:"end_point,omitempty"`
	StopPoint          string    `json:"stop_point,omitempty"`
	StartDateDeparture time.Time `json:"start_date_departure,omitempty"`
	EndDateDeparture   time.Time `json:"end_date_departure,omitempty"`
	PriceMin           int       `json:"price_min,omitempty"`
	PriceMax           int       `json:"price_max,omitempty"`
	TravelTime         string    `json:"travel_time,omitempty"`
	Coupe              bool      `json:"coupe,omitempty"`
	Platzcart          bool      `json:"platzcart,omitempty"`
}
