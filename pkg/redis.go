package pkg

import (
	"github.com/go-redis/redis"
	"strconv"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(address, password string, db string) (*Redis, error) {

	dbInt, err := strconv.Atoi(db)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       dbInt,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Redis{
		Client: client,
	}, nil
}
