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

	c, err := db.GetCacheClientWithDB(db.DB0)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return "", err
	}

	// store shorten url in redis
	if err := c.Set(context.TODO(), shortenURL, inURL,
		time.Second*time.Duration(
			util.GetConfig().TTLShortenKey,
		)).Err(); err != nil {
		util.Logger().Error("error saving shorten url", zap.Error(err))
		return "", err
	}

	c2, err := db.GetCacheClientWithDB(db.DB1)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return "", err
	}
	if err := c2.ZIncrBy(context.TODO(), util.ReqCounterKey, 1, u.Hostname()).Err(); err != nil {
		util.Logger().Error("error increasing counter", zap.Error(err))
		return "", err
	}
	util.Logger().Debug("increased req counter", zap.String("hostname", u.Hostname()))

	return shortenURL, nil
}

func FetchShortenURL(key string) (string, error) {
	c, err := db.GetCacheClientWithDB(db.DB0)
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
	c, err := db.GetCacheClientWithDB(db.DB1)
	if err != nil {
		util.Logger().Error("unable to establish client", zap.Error(err))
		return nil, err
	}

	topDomains, err := c.ZRevRangeWithScores(context.TODO(), util.ReqCounterKey, 0, top-1).Result()
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

func ForTestCreateTestingData() error {
	c2, err := db.GetCacheClientWithDB(db.DB1)
	if err != nil {
		return err
	}

	// prepare some records
	inURLs := []string{
		"https://snghnaveen.1.io/path",
		"https://snghnaveen.2.io/path",
		"https://snghnaveen.3.io/path",
		"https://snghnaveen.4.io/path",
		"https://snghnaveen.5.io/path",
	}

	for _, url := range inURLs {
		if _, err := ShortenURL(url); err != nil {
			return err
		}
	}

	for i := 1; i <= 100; i++ {
		if i%1 == 0 {
			if err := c2.ZIncrBy(context.TODO(),
				util.ReqCounterKey, 1, inURLs[0]).Err(); err != nil {
				return err
			}
		}

		if i%2 == 0 {
			if err := c2.ZIncrBy(context.TODO(),
				util.ReqCounterKey, 1, inURLs[1]).Err(); err != nil {
				return err
			}
		}

		if i%3 == 0 {
			if err := c2.ZIncrBy(context.TODO(),
				util.ReqCounterKey, 1, inURLs[2]).Err(); err != nil {
				return err
			}
		}

		if i%4 == 0 {
			if err := c2.ZIncrBy(context.TODO(),
				util.ReqCounterKey, 1, inURLs[3]).Err(); err != nil {
				return err
			}
		}

		if i%5 == 0 {
			if err := c2.ZIncrBy(context.TODO(),
				util.ReqCounterKey, 1, inURLs[4]).Err(); err != nil {
				return err
			}
		}
	}

	return nil
}
