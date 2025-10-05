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

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:5000",
		Password: "",
		DB:       0})
	resp, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis Client Error : %v", err))
	}
	fmt.Printf("Redis Client Replied: %v", resp)
	storeService.redisClient = redisClient

	return storeService
}

func SaveURLMapping(shortURL, longURL, userID string) {
	err := storeService.redisClient.Set(ctx, shortURL, longURL, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Could not set Long URL << %s >> as Short URL << %s >>.\n ERROR : %v\n", longURL, shortURL, err))
	}
}

func GetLongURL(shortURL string) (longURL string) {
	longURL, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Could not retrieve Long URL from Short URL << %s >>.\n ERROR : %v\n", shortURL, err))
	}
	return
}
