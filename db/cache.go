package db

import (
	"context"
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

func ForTestCreateTestingData() error {
	c, err := GetCacheClientWithDB(DB1)
	if err != nil {
		return err
	}

	// prepare some records
	url1 := "https://snghnaveen.1.io/path"
	url2 := "https://snghnaveen.2.io/path"
	url3 := "https://snghnaveen.3.io/path"
	url4 := "https://snghnaveen.4.io/path"
	url5 := "https://snghnaveen.5.io/path"

	for i := 1; i <= 100; i++ {
		if i%1 == 0 {
			if err := c.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, url1).Err(); err != nil {
				return err
			}
		}

		if i%2 == 0 {
			if err := c.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, url2).Err(); err != nil {
				return err
			}
		}

		if i%3 == 0 {
			if err := c.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, url3).Err(); err != nil {
				return err
			}
		}

		if i%4 == 0 {
			if err := c.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, url4).Err(); err != nil {
				return err
			}
		}

		if i%5 == 0 {
			if err := c.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, url5).Err(); err != nil {
				return err
			}
		}
	}

	return nil
}
