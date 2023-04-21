package db

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/snghnaveen/url-shortener/util"
	"go.uber.org/zap"
)

var cache = make(map[int]*redis.Client)
var mu sync.Mutex

// GetCacheClientWithDB return cache client with requested db
func GetCacheClientWithDB(db int) (*redis.Client, error) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := cache[db]; !ok {
		var rdb *redis.Client

		opt, err := redis.ParseURL(util.GetConfig().RedisURL)
		if err != nil {
			return nil, err
		}

		rdb = redis.NewClient(&redis.Options{
			Addr: opt.Addr,
			DB:   db,
		})

		cache[db] = rdb
	}
	return cache[db], nil
}

func Tmp() {
	ee, err := GetCacheClientWithDB(1)
	util.Logger().Info("tmppppppp", zap.Any("err", err))

	err = ee.Set(context.TODO(),
		"test-key",
		"test-value", time.Second*100).Err()

	util.Logger().Info("tmp22222222", zap.Any("err", err))
	val, err := ee.Get(context.TODO(), "test-key").Result()
	util.Logger().Info("tmp33333", zap.Any("err", err))
	util.Logger().Info("tmp4444444", zap.Any("val", val))

}
