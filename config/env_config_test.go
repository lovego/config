package config

import "fmt"

func ExampleTimestampSign() {
	fmt.Println(TimestampSign(123, "abc"))
	// Output: 21d9bfd0521686c89039b04bf66faf108c391e2334a371dfa51401c5e05a6e32
}
