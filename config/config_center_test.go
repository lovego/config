package config

import (
	"fmt"
	"testing"
)

func TestGetCenterConfig(t *testing.T) {
	config := Get(`../release/img-app/config/test.yml`, `dev`)
	fmt.Println(GetDB(config.Data, `postgres`, `test`))

	c := GetCenterConfig(ConfigCenter{
		Secret:  "123",
		Pull:    config.ConfigCenter.Pull,
		Project: "erp",
		Version: "1.0",
	}, "dev")

	fmt.Println(c)

}
