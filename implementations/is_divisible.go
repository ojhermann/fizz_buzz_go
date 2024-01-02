package implementations

func IsDivisible(dividend int, divisor int) bool {
	if divisor == 0 {
		return false
	}

	return dividend%divisor == 0
}
