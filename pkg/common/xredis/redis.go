package xredis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/conf"
)

var (
	log = logger.Logger("redis")
	cli *RedisClient
)

type RedisClient struct {
	client   *redis.Client
	RedsSync *redsync.Redsync
	Prefix   string
}

func NewRedisClient(cfg conf.Redis) *redis.Client {
	var (
		client   *redis.Client
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// single redis
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Address[0],
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Error(err.Error())
	}
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)

	cli = &RedisClient{client, redsSync, cfg.Prefix}
	return client
}
