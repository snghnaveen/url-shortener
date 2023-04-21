package resolver

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/snghnaveen/url-shortener/util"
	"go.uber.org/zap"
)

// EncodeURL parses the URL string
func EncodeURL(inURL string) (string, error) {
	hash := sha1.New()
	hash.Write([]byte(inURL))
	shortURL := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8]
	util.Logger().Debug("shortURL", zap.String("inURL", shortURL))

	return shortURL, nil
}

// AddScheme will add http in URL
func AddScheme(inURL string) string {
	if inURL[:4] != "http" {
		inURL = "http://" + inURL
	}
	return inURL
}

func FormatURL(inURL string) (*url.URL, error) {
	if !govalidator.IsURL(inURL) {
		return nil, errors.New("invalid url requested")
	}

	inURL = AddScheme(inURL)

	u, err := url.Parse(inURL)
	if err != nil {
		util.Logger().Error("unable to parse error", zap.Error(err))
		return nil, err
	}

	if u.Host == "" {
		util.Logger().Error("invalid url", zap.Error(err))
		return nil, errors.New("invalid url")
	}

	return u, err

	// return inURL, nil
}
