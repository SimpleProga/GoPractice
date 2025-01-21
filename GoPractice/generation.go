package GoPractice

func Generation(a int) string {
	switch {
	case a >= 1946 && a <= 1964:
		return "Hi, boomer"
	case a >= 1965 && a <= 1980:
		return "Hi, X"
	case a >= 1981 && a <= 1996:
		return "Hi, Y"
	case a >= 1997 && a <= 2012:
		return "Hi, Z"
	case a >= 2013:
		return "Hi, Alpha"
	default:
		return "Generation not found"
	}
}
