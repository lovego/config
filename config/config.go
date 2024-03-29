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
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name           string        `yaml:"name" json:"name"`
	RawExternalURL string        `yaml:"externalURL" json:"externalURL"`
	Secret         string        `yaml:"secret" json:"secret"`
	Cookie         Cookie        `yaml:"cookie" json:"cookie"`
	Mailer         string        `yaml:"mailer" json:"mailer"`
	Keepers        []string      `yaml:"keepers" json:"keepers"`
	TimeZone       timeZone      `yaml:"timeZone" json:"timeZone"`
	Data           strmap.StrMap `yaml:"data" json:"data"`

	Env          Environment    `yaml:"-" json:"-"`
	ExternalURL  *url.URL       `yaml:"-" json:"-"`
	TimeLocation *time.Location `yaml:"-" json:"-"`
	path         string

	// 必填字段
	ConfigCenter ConfigCenter `yaml:"configCenter" json:"configCenter" c:"配置中心"`
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

// Get 获取配置时，支持从配置中心获取，也支持从本地文件获取
//  获取本地配置文件中的配置中心地址，
//  如果获取到配置中心地址，则直接从配置中心读取
//  如果没有配置中心地址，则直接从本地文件获取
func Get(path, env string) (config *Config) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	config = &Config{path: path}
	if err := yaml.Unmarshal(content, config); err != nil {
		log.Fatalf("parse %s: %v", path, err)
	}
	defer func() {
		config.init(env)
	}()

	if config.ConfigCenter.Pull == "" {
		return
	}

	u, err := url.Parse(config.ConfigCenter.Pull)
	if err != nil {
		panic(err)
	}

	config.ConfigCenter.Project = u.Query().Get("project")
	config.ConfigCenter.Secret = u.Query().Get("secret")
	config.ConfigCenter.Version = u.Query().Get("version")


	config = GetCenterConfig(config.ConfigCenter, env)
	return
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
	} else {
		c.TimeLocation = time.Local
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
