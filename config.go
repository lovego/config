package config

import (
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/lovego/config/config"
)

var theConfig = config.Get(filepath.Join(Dir(), `config.yml`), Env().Major()).Get(Env().String())

func Name() string {
	return theConfig.Name
}

func DeployName() string {
	return theConfig.DeployName()
}

func ExternalURL() url.URL {
	return theConfig.ExternalURL
}

func Secret() string {
	return theConfig.Secret
}

func Cookie() http.Cookie {
	return theConfig.HttpCookie()
}

func TimestampSign(timestamp int64) string {
	return theConfig.TimestampSign(timestamp)
}

func TimeZone() *time.Location {
	return theConfig.TimeLocation
}

func Keepers() []string {
	return theConfig.Keepers
}
