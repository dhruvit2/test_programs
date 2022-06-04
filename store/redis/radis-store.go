package radisstore

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

func NewStore() (*StorageService, err) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	msg, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nRedis started successfully: message = {%s}", msg)
	return &StorageService{ redisClient: redisClient}, nil
}

// if user id was not provided generate one on the fly : case for not logged in users

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/
func (s *StorageService) SaveUrlMapping(shortUrl string, originalUrl string, userId string) error{
	err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		return errors.New("Failed saving key url | Error: " + err.Error() + " - shortUrl: " + shortUrl + " - originalUrl: " + originalUrl))
	}

	return nil
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func (s *StorageService) RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", errors.New("Failed RetrieveInitialUrl url | Error: " + err.Error() + "- shortUrl: " + shortUrl)
	}

	return result, nil
}