package handler

import (
	"booking_service/configs"
	"booking_service/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) CreateBookedTicket(ctx *gin.Context) {
	// Получение X-Key
	xKey := ctx.GetHeader("X-Key")
	if xKey == "" {
		log.Println("Empty X-Key")
		ctx.JSON(400, Response{
			Error: "Empty X-Key",
		})
		return
	}

	if xKey != "18c690328fb8bbf53a4e5448beb100a035da9191cdea55cb5d67de8f61b646072b66a82db20c07cf2e78293f25e1152bb1a9e749c7622f1dabc6ddc1036ebf74bb18e658714cfb604e543b04f2dfd2d6e3f42a040d3c9cc376d33134fe1b904719d854871a24b8475b77cc0bc1f824881529f5f86351191dc6c1e7449a0b5c18" {
		log.Println("Invalid X-Key")
		ctx.JSON(400, Response{
			Error: "Invalid X-Key",
		})
		return
	}

	// Получение данных из тела запроса
	var ticket entity.BookedTicket

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "Некорретное тело запроса",
		})
		return
	}

	// Валидация данных
	if err := ticket.Validate(); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: err.Error(),
		})
		return
	}

	// Сохранение данных
	_, err := h.services.BookedTicket.Create(ticket)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "invalid request",
		})
		return
	}

	// Отправка ответа
	ctx.JSON(201, Response{
		Message: "OK",
		//Details: id,
	})
	return
}

func (h *Handler) CreateQueueBookedTicket(ctx *gin.Context) {
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
	log.Println(user)

	// Отправка в очередь RabbitMQ
	err = h.services.BookedTicket.SendToRabbitMQ(entity.Income{
		TrainID: 1,
		SeatID:  1,
	}, "auto_booking", configs.RabbitURI)
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
		Details: user,
	})
}
