package testdata

import "fmt"

func ContainLiteral(n int) {
	func(m int) {
		fmt.Println(m % 2)
	}(n / 3)
}
