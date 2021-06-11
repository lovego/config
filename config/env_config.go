package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lovego/duration"
)

type EnvConfig struct {
	Name        string      `yaml:"-"`
	Env         Environment `yaml:"-"`
	ExternalURL url.URL     `yaml:"externalURL"`
	Secret      string      `yaml:"secret"`
	Cookie      Cookie      `yaml:"cookie"`

	Mailer       string   `yaml:"mailer"`
	Keepers      []string `yaml:"keepers"`
	TimeZone     timeZone `yaml:"timeZone"`
	TimeLocation *time.Location
}

// If use http.Cookie, it has no yaml tags, upper camel case is required, so define a new one.
type Cookie struct {
	Name   string            `yaml:"name"`
	Domain string            `yaml:"domain"`
	Path   string            `yaml:"path"`
	MaxAge duration.Duration `yaml:"maxAge"`
}

type timeZone struct {
	Name   string `yaml:"name"`
	Offset int    `yaml:"offset"`
}

func (c *EnvConfig) init(name, env string) {
	c.Name = name
	c.Env = NewEnv(env)
	if c.TimeZone.Name != `` {
		c.TimeLocation = time.FixedZone(c.TimeZone.Name, c.TimeZone.Offset)
	}
}

func (c *EnvConfig) DeployName() string {
	return c.Name + `.` + c.Env.String()
}

func (c *EnvConfig) TimestampSign(ts int64) string {
	return TimestampSign(ts, c.Secret)
}

func TimestampSign(ts int64, secret string) string {
	bytes32 := sha256.Sum256([]byte(fmt.Sprintf("%d,%s", ts, secret)))
	sign := hex.EncodeToString(bytes32[:])
	return sign
}

func (c *EnvConfig) HttpCookie() http.Cookie {
	return http.Cookie{
		Name:   c.Cookie.Name,
		Domain: c.Cookie.Domain,
		Path:   c.Cookie.Path,
		MaxAge: int(c.Cookie.MaxAge.Value / duration.Second),
	}
}
