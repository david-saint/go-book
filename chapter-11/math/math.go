package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	if len(xs) == 0 {
		return float64(0)
	}
	return total / float64(len(xs))
}
