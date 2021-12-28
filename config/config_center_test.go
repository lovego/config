package config

import (
	"fmt"
	"testing"
)

func TestGetCenterConfig(t *testing.T) {
	config := Get(`../release/img-app/config/test.yml`, `qa2`)
	fmt.Println(GetDB(config.Data, `postgres`, `test`))

	c := GetCenterConfig(ConfigCenter{
		Secret:  "123",
		Pull:    config.ConfigCenter.Pull,
		Project: "erp",
		Version: "1.0",
	}, "qa2")

	fmt.Println(c)

}
