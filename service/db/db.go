package db

import (
	"github.com/go-redis/redis"
)

func NewDatabase(address string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
