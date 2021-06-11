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

func ExampleExternalURL() {
	fmt.Println(ExternalURL().String())

	// Output: https://example.com/home
}
