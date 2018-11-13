package testdata

func literal() func(int) bool {
	return func(n int) bool {
		if n%2 == 0 {
			return false
		}
		return true
	}
}
