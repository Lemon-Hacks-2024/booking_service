package handler

import (
	"booking_service/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func (h *Handler) GetTrains(ctx *gin.Context) {
	var input entity.TrainsInputQueryParam
	// Получение данных из тела запроса
	input.BookingAvailable = ctx.Query("booking_available")
	input.StartPoint = ctx.Query("start_point")
	input.EndPoint = ctx.Query("end_point")
	input.StartDateDeparture, _ = time.Parse("02.01.2006 15:04:05", ctx.Query("start_date_departure"))
	input.EndDateDeparture, _ = time.Parse("02.01.2006 15:04:05", ctx.Query("end_date_departure"))
	input.TravelTime = ctx.Query("travel_time")
	// Валидация данных
	// TODO:

	// Получение данных
	trains, err := h.services.Train.GetAllTrains(input)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "invalid request",
		})
		return
	}
	// Отправка ответа
	ctx.JSON(200, Response{
		Message: "OK",
		Details: trains,
	})
	return
}
