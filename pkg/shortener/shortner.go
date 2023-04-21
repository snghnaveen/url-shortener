package shortener

import (
	"context"
	"time"

	"github.com/snghnaveen/url-shortener/db"
	"github.com/snghnaveen/url-shortener/pkg/resolver"
	"github.com/snghnaveen/url-shortener/pkg/rest"
	"github.com/snghnaveen/url-shortener/util"
	"go.uber.org/zap"
)

const (
	DB0 = iota
	DB1

	reqCounterKey = "req_counter"
)

func ShortenURL(inURL string) (string, error) {
	util.Logger().Debug("requested inURL", zap.String("inURL", inURL))
	u, err := resolver.FormatURL(inURL)
	if err != nil {
		return "", err
	}
	util.Logger().Debug("formatted inURL", zap.String("u.String", u.String()))

	shortenURL, err := resolver.EncodeURL(u.String())
	if err != nil {
		return "", err
	}
	util.Logger().Info("calculated shorten url", zap.String("shortenUrl", shortenURL))

	c, err := db.GetCacheClientWithDB(DB0)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return "", err
	}

	// store shorten url in redis
	if err := c.Set(context.TODO(), shortenURL, inURL, time.Second*100).
		Err(); err != nil {
		util.Logger().Error("error saving shorten url", zap.Error(err))
		return "", err
	}

	c2, err := db.GetCacheClientWithDB(DB1)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return "", err
	}
	if err := c2.ZIncrBy(context.TODO(), reqCounterKey, 1, u.Hostname()).Err(); err != nil {
		util.Logger().Error("error increasing counter", zap.Error(err))
		return "", err
	}
	util.Logger().Debug("increased req counter", zap.String("hostname", u.Hostname()))

	return shortenURL, nil
}

func FetchShortenURL(key string) (string, error) {
	c, err := db.GetCacheClientWithDB(DB0)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return "", err
	}

	util.Logger().Info("requested fetch shorten url key", zap.String("key", key))
	val, err := c.Get(context.TODO(), key).Result()
	if err != nil {
		return "", rest.ErrShortenURLNotFound
	}
	return val, nil
}

func GetTopRequested(top int64) ([]map[string]interface{}, error) {
	c, err := db.GetCacheClientWithDB(DB1)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return nil, err
	}

	topDomains, err := c.ZRevRangeWithScores(context.TODO(), reqCounterKey, 0, top-1).Result()
	if err != nil {
		util.Logger().Error("unable to fetch ZRevRangeWithScores", zap.Error(err))
		return nil, err
	}

	out := make([]map[string]interface{}, 0)
	for i, record := range topDomains {
		out = append(out, map[string]interface{}{
			"rank":  i + 1,
			"url":   record.Member,
			"score": record.Score,
		})

		util.Logger().Info("top requested domain ranking",
			zap.Any("rank", i+1),
			zap.Any("url", record.Member),
			zap.Any("score", record.Score),
		)
	}
	return out, err
}
