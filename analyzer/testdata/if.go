package testdata

func oneCondition(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func elseif(n int) bool {
	if n%2 == 0 {
		return true
	} else if n%3 == 0 {
		return true
	}
	return false
}
