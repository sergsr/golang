package solutions

func myPow(x float64, n int) float64 {
	result := 1.0
	accumulator := x
	neg := n < 0
	bits := n
	if neg {
		bits = -1 * n
	}
	for bits > 0 {
		if bits%2 != 0 {
			result *= accumulator
		}
		bits = bits >> 1
		accumulator *= accumulator
	}
	if neg {
		result = 1.0 / result
	}
	return result
}
