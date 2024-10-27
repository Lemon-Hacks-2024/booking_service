package handler

import (
	"booking_service/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes(port string) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Настройка CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		//
		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}

		c.Next()
	})
	//
	api := router.Group("/ax-train")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})

		api.POST("/register", h.RegisterUser)
		api.POST("/login", h.LoginUser)

		trains := api.Group("/trains")
		{
			trains.GET("/", h.GetTrains)
		}

		wagons := api.Group("/wagons")
		{
			wagons.GET("/", h.GetWagons)
		}

		users := api.Group("/users")
		{
			users.GET("/", h.GetUser)
		}

		tickets := api.Group("/booked-tickets")
		{
			tickets.POST("/", h.CreateBookedTicket)
		}

		mocks := api.Group("/mocks")
		{
			mocks.GET("/api/info/trains", h.MockTrains)
			mocks.GET("/api/info/seats", h.MockSeats)
		}

	}

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
