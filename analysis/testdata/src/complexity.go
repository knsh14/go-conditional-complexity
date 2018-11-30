package main

func GetOrdinalSuffix(n int) string { // NG - testdata/src/complexity/complexity.go:3 func GetOrdinalSuffix complexity=5
	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
