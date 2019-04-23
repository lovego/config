package postgres

import (
	"fmt"
	"log"
)

func ExampleDB() {
	db := DB("test")
	var v int64
	if err := db.Query(&v, "SELECT 2019"); err != nil {
		log.Panic(v)
	}
	fmt.Println(v)

	// Output: 2019
}
