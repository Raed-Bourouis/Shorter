package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDuration = 6 * time.Hour

func initStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:5000", Password: "", DB: 0})
	resp, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis Client Error : %v", err))
	}
	fmt.Printf("Redis Client Replied: %v", resp)
	storeService.redisClient = redisClient

	return storeService
}
