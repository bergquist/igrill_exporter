package main

func mean(vals []float64) float64 {
	total := 0.0
	for _, val := range vals {
		total += float64(val)
	}
	return total / float64(len(vals))
}
