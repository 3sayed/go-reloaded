package formulas


func Variance(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	mean := Average(numbers)
	var sum float64
	for _, number := range numbers {
		diff := number - mean
		sum += diff * diff
	}
	return sum / float64(len(numbers))
}