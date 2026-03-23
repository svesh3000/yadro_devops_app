package application

import "sort"

func calcAverage(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range vals {
		sum += v
	}

	return sum / float64(len(vals))
}

func calcMedian(vals []float64) float64 {
	n := len(vals)
	if n == 0 {
		return 0
	}

	sorted := make([]float64, n)
	copy(sorted, vals)
	sort.Float64s(sorted)

	if n%2 != 0 {
		return sorted[n/2]
	}

	return (sorted[n/2-1] + sorted[n/2]) / 2
}
