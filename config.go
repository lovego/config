package config

import (
	"net/http"
	"net/url"
	"time"

	"github.com/lovego/config/config"
)

var theConfig = config.Get(FilePath(Env()), Env().String())

func Name() string {
	return theConfig.Name
}

func DeployName() string {
	return theConfig.DeployName()
}

func ExternalURL() *url.URL {
	return theConfig.ExternalURL
}

func ExternalURLIsHTTPS() bool {
	return theConfig.ExternalURLIsHTTPS()
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
