package conf

import "fmt"

func ExampleGetDb() {
	strMap := Data(`../release/img-app/config/envs/test.yml`)
	fmt.Println(GetDb(strMap, `postgres`, `test`))

	if v, err := GetDb(strMap, `postgres`, `shards`); err != nil {
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
	// postgres://postgres:@localhost/test?sslmode=disable <nil>
	// {No:1 Url:postgres://postgres:@localhost/test_1?sslmode=disable}
	// {No:2 Url:postgres://postgres:@localhost/test_2?sslmode=disable}
	// {IdSeqIncrementBy:1000}
}
