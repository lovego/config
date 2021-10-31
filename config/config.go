// The parent config package get root dir when init, it panics if no root is found.
// So this package extract the common code, so it can be used by xiaomei.
package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/lovego/duration"
	"github.com/lovego/strmap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name           string        `yaml:"name"`
	RawExternalURL string        `yaml:"externalURL"`
	Secret         string        `yaml:"secret"`
	Cookie         Cookie        `yaml:"cookie"`
	Mailer         string        `yaml:"mailer"`
	Keepers        []string      `yaml:"keepers"`
	TimeZone       timeZone      `yaml:"timeZone"`
	Data           strmap.StrMap `yaml:"data"`

	Env          Environment    `yaml:"-"`
	ExternalURL  *url.URL       `yaml:"-"`
	TimeLocation *time.Location `yaml:"-"`
	path         string
}

// If use http.Cookie, it has no yaml tags, upper camel case is required, so define a new one.
type Cookie struct {
	Name     string            `yaml:"name"`
	Domain   string            `yaml:"domain"`
	Path     string            `yaml:"path"`
	MaxAge   duration.Duration `yaml:"maxAge"`
	Secure   bool              `yaml:"secure"`
	HttpOnly bool              `yaml:"httpOnly"`
	SameSite string            `yaml:"sameSite"`
}

type timeZone struct {
	Name   string `yaml:"name"`
	Offset int    `yaml:"offset"`
}

func Get(path, env string) *Config {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{path: path}
	if err := yaml.Unmarshal(content, config); err != nil {
		log.Fatalf("parse %s: %v", path, err)
	}

	config.init(env)
	return config
}

func (c *Config) init(env string) {
	c.Env = NewEnv(env)
	if c.RawExternalURL != `` {
		if u, err := url.Parse(c.RawExternalURL); err != nil {
			log.Fatalf("parse externalURL: %v", err)
		} else {
			c.ExternalURL = u
		}
	}
	if c.TimeZone.Name != `` {
		c.TimeLocation = time.FixedZone(c.TimeZone.Name, c.TimeZone.Offset)
	}
	switch c.Cookie.SameSite {
	case "", "lax", "strict", "none":
	default:
		log.Fatalf("invalid sameSite: %s", c.Cookie.SameSite)
	}
}

func (c *Config) DeployName() string {
	return c.Name + `.` + c.Env.String()
}

func (c *Config) ExternalURLIsHTTPS() bool {
	return c.ExternalURL.Scheme == "https" // url.Parse always strings.toLower schema.
}

func (c *Config) TimestampSign(ts int64) string {
	return TimestampSign(ts, c.Secret)
}

func TimestampSign(ts int64, secret string) string {
	bytes32 := sha256.Sum256([]byte(fmt.Sprintf("%d,%s", ts, secret)))
	sign := hex.EncodeToString(bytes32[:])
	return sign
}

func (c *Config) HttpCookie() http.Cookie {
	cookie := http.Cookie{
		Name:     c.Cookie.Name,
		Domain:   c.Cookie.Domain,
		Path:     c.Cookie.Path,
		MaxAge:   int(c.Cookie.MaxAge.Value / duration.Second),
		Secure:   c.Cookie.Secure,
		HttpOnly: c.Cookie.HttpOnly,
	}

	setSameSiteMode(&cookie, c.Cookie.SameSite)
	return cookie
}
