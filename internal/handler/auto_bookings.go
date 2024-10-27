package handler

import (
	"booking_service/configs"
	"booking_service/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Bookings struct {
	ID             int      `json:"id"`
	Startpoint     string   `json:"startpoint"`
	Endpoint       string   `json:"endpoint"`
	WagonType      string   `json:"wagon_type"`
	TicketCount    int      `json:"ticket_count"`
	SeatPreference []string `json:"seat_preference"`
	DepartureDates []string `json:"departure_dates"`
	TrainID        int      `json:"train_id"`
	WagonID        int      `json:"wagon_id"`
	SeatID         int      `json:"seat_id"`
}

func (h *Handler) AutoBookings(ctx *gin.Context) {
	// Получения токена из заголовков запроса
	token := ctx.GetHeader("Authorization")
	if token == "" {
		log.Println("empty token")
		ctx.JSON(400, Response{
			Error: "empty token",
		})
		return
	}

	// Проверка токена
	user, err := h.services.User.GetUserByToken(token)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "invalid token",
		})
		return
	}

	var bookings Bookings
	if err := ctx.ShouldBindJSON(&bookings); err != nil {
		ctx.JSON(400, Response{
			Error: "Некорретное тело запроса",
		})
		return
	}
	dateFormat := "02.01.2006"
	dateFrom, _ := time.Parse(dateFormat, bookings.DepartureDates[0])
	dateFrom = time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, dateFrom.Location())

	var dateTo time.Time
	if len(bookings.DepartureDates) > 1 {
		dateTo, _ = time.Parse(dateFormat, bookings.DepartureDates[len(bookings.DepartureDates)-1])
	} else {
		dateTo, _ = time.Parse(dateFormat, bookings.DepartureDates[0])
	}
	dateTo = time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 23, 59, 59, 0, dateTo.Location())

	log.Println(dateFrom)
	log.Println(dateFrom)

	income := entity.Income{
		UserID:        user.ID,
		TrainID:       bookings.TrainID,
		WagonID:       bookings.WagonID,
		SeatID:        bookings.SeatID,
		Route:         fmt.Sprintf("%s -> %s", bookings.Startpoint, bookings.Endpoint),
		DateFrom:      dateFrom.Format("02.01.2006 15:04:05"),
		DateTo:        dateTo.Format("02.01.2006 15:04:05"),
		WagonType:     bookings.WagonType,
		PlacePosition: bookings.SeatPreference,
		Price:         0,
		SeatsQty:      bookings.TicketCount,
		NeedNearby:    false,
	}

	// Отправка в RabbitMQ
	err = h.services.BookedTicket.SendToRabbitMQ(income, "auto_booking", configs.RabbitURI)
	if err != nil {
		ctx.JSON(400, Response{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(200, Response{
		Message: "OK",
	})
}
