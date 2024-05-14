package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// ConnectToRedis initialise la connexion à la base de données Redis
func ConnectToRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	address := fmt.Sprintf("%s:%s", host, port)
	// Connexion au serveur Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	// Vérifier la connexion
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Impossible de se connecter à Redis: %v", err)
	}

	return rdb
}
