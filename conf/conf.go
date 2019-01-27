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

	Mailer       string   `yaml:"mailer"`
	Keepers      []string `yaml:"keepers"`
	TimeZone     timeZone `yaml:"timeZone"`
	TimeLocation *time.Location
}

type timeZone struct {
	Name   string `yaml:"name"`
	Offset int    `yaml:"offset"`
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
