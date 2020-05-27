package conf

import "fmt"

func ExampleParseDuation() {
	fmt.Println(ParseDuration(`1h3m5`))
	fmt.Println(ParseDuration(`1h3m5`))

	// Output:
	// 3785 <nil>
	// 3785 <nil>
}

func ExampleParseDuation_pure() {
	fmt.Println(ParseDuration(``))
	fmt.Println(ParseDuration(`0`))
	fmt.Println(ParseDuration(`1`))
	fmt.Println(ParseDuration(`12`))
	fmt.Println(ParseDuration(`123`))

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 12 <nil>
	// 123 <nil>
}

func ExampleParseDuation_seconds() {
	fmt.Println(ParseDuration(`0s`))
	fmt.Println(ParseDuration(`1s`))
	fmt.Println(ParseDuration(`12s`))
	fmt.Println(ParseDuration(`123s`))

	// Output:
	// 0 <nil>
	// 1 <nil>
	// 12 <nil>
	// 123 <nil>
}

func ExampleParseDuation_minutes() {
	fmt.Println(ParseDuration(`0m`))
	fmt.Println(ParseDuration(`1m`))
	fmt.Println(ParseDuration(`12m`))
	fmt.Println(ParseDuration(`123m`))

	// Output:
	// 0 <nil>
	// 60 <nil>
	// 720 <nil>
	// 7380 <nil>
}
