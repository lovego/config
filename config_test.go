package config

import (
	"fmt"
	"strings"
)

func ExampleRoot() {
	root := Root()
	fmt.Println(strings.HasSuffix(root, "/config/release/img-app"))

	// Output: true
}
