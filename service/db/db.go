package db

import (
	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

//type Test1 struct {
//	Key   string
//	Value int
//}

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	//json := `{"test": 22}`
	//json1 := `{"gag": 23}`
	//err := client.Set("test", json, 0).Err()
	//err1 := client.Set("id1", json1, 0).Err()

	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if err1 != nil {
	//	fmt.Println(err)
	//}
	//
	//val, err := client.Get("id1").Result()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(val)

	return &Database{
		Client: client,
	}, nil
}
