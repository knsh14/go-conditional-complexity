package testdata

func simpleSwitch(n int) {
	switch n {
	case 0:
		println("zero")
	}
}

func multiCase(n int) {
	switch n {
	case 0:
		println("zero")
	case 1:
		println("zero")
	}
}

func twoMatchInOneCase(n int) {
	switch {
	case 0 < n, n%7 == 0:
		println("foo")
	}
}

func withDefault(n int) {
	switch n {
	case 0:
		println("zero")
	default:
		println("non zero")
	}
}
