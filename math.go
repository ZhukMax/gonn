package gonn

import "math"

/**
Math functions
 */

func Sigmoid(x float64) float64 {
	var output float64
	a := -1 * x
	output = 1 / (1 + math.Pow(math.E, a))

	return output
}

func Delta()  {
	// todo
}

// Delta for output layer's nodes
func DeltaOut(x, i float64) float64 {
	return (i - x) * ((1 - x) * x)
}
