package main

func max(a []float64) float64 {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func argmax(a []float64) int {
	max := a[0]
	max_index := 0
	for i, value := range a {
		if value > max {
			max = value
			max_index = i
		}
	}
	return max_index
}
