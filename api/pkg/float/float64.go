package float

type Float struct {
}

func Compare(a, b float64) int {
	if a < b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}
