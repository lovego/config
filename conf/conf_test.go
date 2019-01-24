package conf

import "fmt"

func ExampleTimestampSign() {
	ts, sign := TimestampSign("abc")
	fmt.Println(ts > 0, len(sign))
	// Output: true 64
}
