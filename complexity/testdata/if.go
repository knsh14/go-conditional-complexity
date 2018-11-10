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

func and(n int) bool {
	if n%2 == 0 && n < 0 {
		return true
	}
	return false
}

func or(n int) bool {
	if n%2 == 0 || n < 0 {
		return true
	}
	return false
}

// assign is one complexity with assign and check statements.
// TODO: fix?
func assign(n interface{}) bool {
	if v, ok := n.(uint64); ok {
		println(v)
		return true
	}
	return false
}
