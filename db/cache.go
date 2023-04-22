package db

import (
	"sync"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/snghnaveen/url-shortener/util"
	"go.uber.org/zap"
)

const (
	DB0 = iota
	DB1
)

var cache = make(map[int]*redis.Client)
var mu sync.Mutex
var once sync.Once

// GetCacheClientWithDB return cache client with requested db
func GetCacheClientWithDB(db int) (*redis.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := cache[db]; !ok {
		var rdb *redis.Client
		if util.IsTestingMode() {
			util.Logger().Info("running testing mode redis",
				zap.String("ForTestGetRedisURL", ForTestGetRedisURL()))
			rdb = redis.NewClient(&redis.Options{
				Addr: ForTestGetRedisURL(),
				DB:   db,
			})
		} else {
			opt, err := redis.ParseURL(util.GetConfig().RedisURL)
			if err != nil {
				return nil, err
			}

			rdb = redis.NewClient(&redis.Options{
				Addr: opt.Addr,
				DB:   db,
			})
		}

		cache[db] = rdb
	}
	return cache[db], nil
}

var testClientUrl string

func ForTestGetRedisURL() string {
	once.Do(func() {
		s, err := miniredis.Run()
		if err != nil {
			panic("unable to run miniredis")
		}
		testClientUrl = s.Addr()
	})
	return testClientUrl
}
