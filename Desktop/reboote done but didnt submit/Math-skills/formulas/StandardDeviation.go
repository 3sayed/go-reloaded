package formulas

import(
	"math"
)

func StndardDeviation(numbers []float64) float64 {
	
	if len(numbers) == 0 {
		return 0
	}

	variance := Variance(numbers)

	return math.Sqrt(variance)
}