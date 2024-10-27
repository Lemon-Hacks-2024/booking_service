package handler

import (
	"booking_service/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	// Получение данных из тела запроса
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "Некорретное тело запроса",
		})
		return
	}

	// Валидация данных
	if err := input.ValidateUserByRegistration(); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: err.Error(),
		})
		return
	}

	// Регистрация пользователя
	token, err := h.services.User.CreateUser(input)
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
		Details: token,
	})
	return
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	// Получение данных из тела запроса
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: "Некорретное тело запроса",
		})
		return
	}

	// Валидация данных
	if err := input.ValidateUserByLogin(); err != nil {
		log.Println(err)
		ctx.JSON(400, Response{
			Error: err.Error(),
		})
		return
	}

	// Авторизация пользователя
	token, err := h.services.User.Login(input.Email, input.Password)
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
		Details: token,
	})
	return
}

func (h *Handler) GetUser(ctx *gin.Context) {
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

	// Отправка ответа
	ctx.JSON(200, Response{
		Message: "OK",
		Details: entity.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	})
	return
}
