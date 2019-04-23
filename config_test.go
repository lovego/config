package config

import (
	"fmt"
	"strings"
)

func ExampleRoot() {
	root := Root()
	fmt.Println(strings.HasSuffix(root, "/github.com/lovego/config/release/img-app"))

	// Output: true
}
