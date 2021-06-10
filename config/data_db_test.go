package config

import "fmt"

func ExampleGetDB() {
	strMap := Data(`../release/img-app/config/envs/test.yml`)
	fmt.Println(GetDB(strMap, `postgres`, `test`))

	if v, err := GetDB(strMap, `postgres`, `shards`); err != nil {
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
