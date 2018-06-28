package dbclient

import "github.com/perennial-microservices/blog/accountservice/model"
import (
	"github.com/go-redis/redis"
	"strconv"
	"fmt"
)

type IRedisClient interface {
	StartRedis()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
	Check() (bool)
}

type RedisClient struct {
	client *redis.Client
}

func (rc *RedisClient) StartRedis() {
	rc.client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
}

func (rc *RedisClient) QueryAccount(accountId string) (model.Account, error) {
	name, err := rc.client.Get(accountId).Result()

	if err != nil {
		return model.Account{}, err
	}

	account := model.Account{
		Id:accountId,
		Name:name,
	}

	return account, nil
}

func (rc *RedisClient) Seed() {
	total := 100
	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000+i)

		acc := model.Account{
			Id: key,
			Name: "Person_" + strconv.Itoa(i),
		}

		_ = rc.client.Set(acc.Id, acc.Name, 0).Err()
	}

	fmt.Printf("Seeded %v fake accounts...\n", total)
}

func (rc *RedisClient) Check() (bool) {
	_, err := rc.client.Ping().Result()

	return err == nil
}