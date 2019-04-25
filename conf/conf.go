package conf

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Conf struct {
	Name   string `yaml:"-"`
	Env    string `yaml:"-"`
	Https  bool   `yaml:"https"`
	Domain string `yaml:"domain"`
	Secret string `yaml:"secret"`
	Cookie Cookie `yaml:"cookie"`

	Mailer       string   `yaml:"mailer"`
	Keepers      []string `yaml:"keepers"`
	TimeZone     timeZone `yaml:"timeZone"`
	TimeLocation *time.Location
}

type Cookie struct {
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
	Path   string `yaml:"path"`
	MaxAge int    `yaml:"maxAge"`
}

type timeZone struct {
	Name   string `yaml:"name"`
	Offset int    `yaml:"offset"`
}

func (c *Conf) init(name, env string) {
	c.Name = name
	c.Env = env
	if c.TimeZone.Name != `` {
		c.TimeLocation = time.FixedZone(c.TimeZone.Name, c.TimeZone.Offset)
	}
}

func (c *Conf) DeployName() string {
	return c.Name + `-` + c.Env
}

func (c *Conf) Url() string {
	if c.Https {
		return "https://" + c.Domain
	} else {
		return "http://" + c.Domain
	}
}

func (c *Conf) TimestampSign(ts int64) string {
	return TimestampSign(ts, c.Secret)
}

func TimestampSign(ts int64, secret string) string {
	bytes32 := sha256.Sum256([]byte(fmt.Sprintf("%d,%s", ts, secret)))
	sign := hex.EncodeToString(bytes32[:])
	return sign
}
