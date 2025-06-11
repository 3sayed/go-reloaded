package formulas

func Median(numbers [] float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	n := len(numbers)
	if n%2 == 1 {
		return numbers[n/2]
	} else{
		return (numbers[n/2-1] + numbers[n/2]) / 2.0
	}
}