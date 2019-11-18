package semver_test

import (
	"fmt"

	"github.com/kinbiko/semver"
)

func ExampleUpversion() {
	got, _ := semver.Upversion("minor", "3.2.6")
	fmt.Println(got)
	//output:
	//3.3.0
}

func ExampleUpversion_withVPrefix() {
	got, _ := semver.Upversion("minor", "v3.2.6")
	fmt.Println(got)
	//output:
	//v3.3.0
}
