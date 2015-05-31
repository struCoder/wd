package wd

import (
	"fmt"
)

func ExampleNumberLines() {
	s := `line 1
line 2
line 3`
	fmt.Println(NumberLines(s))
	// Output:
	// 1 | line 1
	// 2 | line 2
	// 3 | line 3
}
