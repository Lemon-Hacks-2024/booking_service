package service

import (
	"booking_service/internal/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var AxTrainToken = "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NUBxd2UucnUiLCJpYXQiOjE3Mjk5NDYxMTIsImV4cCI6MTczMDAzMjUxMn0.a-iDnW5H1JCXoHaVf8wXlfwGU7r_VtXgGEPpRj6D6-c"

type AxTrainService struct {
}

func NewAxTrainService() *AxTrainService {
	return &AxTrainService{}
}

func (s *AxTrainService) GetWagonsByTrain(trainID int) ([]entity.Wagon, error) {
	// Захватываем "слот" для выполнения запроса
	//requestLimit <- struct{}{}
	//defer func() {
	//	<-requestLimit                     // Освобождаем "слот" после завершения
	//	time.Sleep(100 * time.Millisecond) // Добавляем задержку в 200 миллисекунд
	//}()

	var wagons []entity.Wagon

	pathUrl := fmt.Sprintf("http://84.252.135.231/api/info/wagons?trainId=%d", trainID)

	req, err := http.NewRequest("GET", pathUrl, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", AxTrainToken)
	client := &http.Client{
		//Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var wagonsResponse []entity.WagonResponse
	err = json.Unmarshal(body, &wagonsResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, wagon := range wagonsResponse {
		seats := make([]entity.Seat, 0)

		for _, seat := range wagon.Seats {

			seatNum, err := strconv.Atoi(seat.SeatNum)
			if err != nil {
				log.Println(err)
				seatNum = 0
			}

			seatBlock, err := strconv.Atoi(seat.Block)
			if err != nil {
				log.Println(err)
				seatBlock = 0
			}

			seats = append(seats, entity.Seat{
				SeatID:        seat.SeatID,
				SeatNum:       seatNum,
				Block:         seatBlock,
				Price:         seat.Price,
				BookingStatus: seat.BookingStatus,
			})
		}

		wagons = append(wagons, entity.Wagon{
			WagonID: wagon.WagonID,
			Type:    wagon.Type,
			Seats:   seats,
		})
	}

	return wagons, nil
}

// GetAllTrains Поиск всех возможных поездов с различными параметрами
//var requestLimit = make(chan struct{}, 2) // Канал с буфером размером 1

func (s *AxTrainService) GetAllTrains(queryParam entity.TrainsInputQueryParam) ([]entity.Train, error) {
	//// Захватываем "слот" для выполнения запроса
	//requestLimit <- struct{}{}
	//defer func() {
	//	<-requestLimit                     // Освобождаем "слот" после завершения
	//	time.Sleep(100 * time.Millisecond) // Добавляем задержку в 200 миллисекунд
	//}()

	var trains entity.TrainsResponse

	stopPoint := ""
	if queryParam.StopPoint != "" {
		stopPoint += fmt.Sprintf("&stop_point=%s", url.QueryEscape(queryParam.StopPoint))
	}

	url := fmt.Sprintf("http://84.252.135.231/api/info/trains?booking_available=%s&start_point=%s&end_point=%s", queryParam.BookingAvailable, url.QueryEscape(queryParam.StartPoint), url.QueryEscape(queryParam.EndPoint))
	url += stopPoint

	log.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Authorization", AxTrainToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{
		//Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	if resp.StatusCode == 403 {
		token, err := s.Login()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		AxTrainToken = token
		return s.GetAllTrains(queryParam)
	}

	if resp.StatusCode == 429 {
		log.Println(err)
		time.Sleep(200 * time.Millisecond)
		return s.GetAllTrains(queryParam)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(string(respBody))

	err = json.Unmarshal(respBody, &trains)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var trainsByFilter []entity.Train
	for i := 0; i < len(trains); i++ {

		// Фильтрация по дате
		parsedDateDeparture, _ := time.Parse("02.01.2006 15:04:05", trains[i].DetailedRoute[0].Departure)
		if parsedDateDeparture.Format("02.01.2006") < queryParam.StartDateDeparture.Format("02.01.2006") ||
			parsedDateDeparture.Format("02.01.2006") > queryParam.EndDateDeparture.Format("02.01.2006") {
			continue
		}

		// Фильтрация по времени в пути
		if queryParam.TravelTime != "" {
			parsedDateArrival, _ := time.Parse("02.01.2006 15:04:05", trains[i].DetailedRoute[len(trains[i].DetailedRoute)-1].Arrival)
			log.Println("Arrival Time: ", parsedDateArrival)
			// Переводим дату в минуты
			travelTime := int(parsedDateArrival.Sub(parsedDateDeparture).Minutes())
			log.Println("Duration Travel Time: ", travelTime)
			queryParamTravelTime, err := strconv.Atoi(queryParam.TravelTime)
			if err != nil {
				log.Println(err)
			}
			if travelTime >= queryParamTravelTime {
				continue
			}
		}

		var wagonsInfo []entity.TrainsWagonsInfo
		for _, wagon := range trains[i].WagonsInfo {
			wagonsInfo = append(wagonsInfo, entity.TrainsWagonsInfo{
				WagonID: wagon.WagonID,
				Type:    wagon.Type,
			})
		}
		log.Println(trains[i].AvailableSeatsCount)

		// Добавляем объект с данными поезда
		trainsByFilter = append(trainsByFilter, entity.Train{
			TrainID:             trains[i].TrainID,
			Route:               trains[i].GlobalRoute,
			StartPoint:          trains[i].DetailedRoute[0].Name,
			StartPointDeparture: trains[i].StartPointDeparture,
			EndPoint:            trains[i].DetailedRoute[len(trains[i].DetailedRoute)-1].Name,
			EndPointArrival:     trains[i].EndpointArrival,
			WagonsInfo:          wagonsInfo,
			AvailableSeatsCount: trains[i].AvailableSeatsCount,
		})

		trainsByFilter[i].FillFields()
	}

	return trainsByFilter, nil
}

func (s *AxTrainService) Login() (string, error) {

	url := fmt.Sprintf("http://84.252.135.231/api/auth/login")
	log.Println(url)

	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "semyon.albeev@gmail.com",
		Password: "H=8pGjPVR;ErCS,F",
	}
	log.Println(data)
	// Кодирование данных в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Забираем только токен
	var token struct {
		Token string `json:"token"`
	}
	err = json.Unmarshal(respBody, &token)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token.Token, nil
}
