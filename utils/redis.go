package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/saleh-ghazimoradi/GoGarage/logger"
	"time"
)

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
}

func RedisConnection(cfg RedisConfig) (*redis.Client, error) {
	logger.Logger.Info("Connection to redis established")
	re := redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:           cfg.Password,
		DB:                 cfg.Db,
		DialTimeout:        cfg.DialTimeout,
		ReadTimeout:        cfg.ReadTimeout,
		WriteTimeout:       cfg.WriteTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
		PoolSize:           cfg.PoolSize,
		PoolTimeout:        cfg.PoolTimeout,
		IdleTimeout:        cfg.IdleTimeout,
	})

	if err := re.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return re, nil
}
