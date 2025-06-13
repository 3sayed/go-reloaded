package formulas

func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}