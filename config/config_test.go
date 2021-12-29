package config

import (
	"fmt"
	"testing"
)

func ExampleGet() {
	config := Get(`../release/img-app/config/test.yml`, `dev`)
	fmt.Println(GetDB(config.Data, `postgres`, `test`))

	if v, err := GetDB(config.Data, `postgres`, `shards`); err != nil {
		fmt.Println(err)
	} else if shards, ok := v.(*Shards); ok {
		for _, row := range shards.Shards {
			fmt.Printf("%+v\n", row)
		}
		fmt.Printf("%+v\n", shards.Settings)
	} else {
		fmt.Println(v)
	}
	// Output:
	// postgres://postgres:postgres@localhost/postgres?sslmode=disable <nil>
	// {No:1 Url:postgres://postgres:@localhost/test_1?sslmode=disable}
	// {No:2 Url:postgres://postgres:@localhost/test_2?sslmode=disable}
	// {IdSeqIncrementBy:1000}
}

func ExampleTimestampSign() {
	fmt.Println(TimestampSign(123, "abc"))
	// Output: 21d9bfd0521686c89039b04bf66faf108c391e2334a371dfa51401c5e05a6e32
}

func TestMap(t *testing.T) {
	data:= map[string]map[string]string{
		"lch":{
			"lch":"lchjczw",
			"lch1":"lchjczw",
		},
	}


	lch:=data["lch"]

	lch["lch"] = "lch"

	fmt.Println(data)

}