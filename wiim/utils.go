package wiim

func limitValue(value, minimum, maximum int) int {
	return max(minimum, min(maximum, value))
}
