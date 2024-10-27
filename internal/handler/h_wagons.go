package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetWagons(c *gin.Context) {
	// Получение ID поезда из параметров запроса
	trainID, err := strconv.Atoi(c.Query("train_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "неверный формат train_id",
		})
		return
	}

	// Вызов метода GetWagons
	trains, err := h.services.Train.GetWagonsByTrain(trainID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "ошибка получения данных поезда",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"details": trains,
	})
	return
}
