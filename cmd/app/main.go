package main

import (
	"booking_service/configs"
	"booking_service/internal/handler"
	"booking_service/internal/repository"
	"booking_service/internal/service"
	"booking_service/pkg"
	"github.com/joho/godotenv"

	"log"
)

//	@title		Booking Service
//	@version	1.0

//	@BasePath	/api

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	redis, err := pkg.NewRedis(cfg.RedisHost+":"+cfg.RedisPort, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("Ошибка при инициализации подключения к Redis: %s", err.Error())
	}

	db, err := pkg.NewPostgresDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBSSLMode)
	if err != nil {
		log.Fatalf("Ошибка при инициализации подключения к postgres: %s", err.Error())
	}

	repos := repository.NewRepository(db, redis)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(cfg.AppPort)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}
