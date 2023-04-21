package shortener

import (
	"context"
	"time"

	"github.com/snghnaveen/url-shortner/db"
	"github.com/snghnaveen/url-shortner/pkg/resolver"
	"github.com/snghnaveen/url-shortner/pkg/rest"
	"github.com/snghnaveen/url-shortner/util"
	"go.uber.org/zap"
)

const (
	DB0 = iota
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
