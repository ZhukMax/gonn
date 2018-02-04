package gonn

import (
	"math"
)

/**
Math functions
 */

 /**
 @public
  */
func Sigmoid(x float64) float64 {
	var output float64
	a := -1 * x
	output = 1 / (1 + math.Pow(math.E, a))

	return output
}

/**
Delta for hidden layer's nodes
@public
 */
func Delta(out, amount float64) float64 {
	return f(out) * amount
}

/**
Delta for output layer's nodes
@public
 */
func DeltaOut(out, ideal float64) float64 {
	return (ideal - out) * f(out)
}

/**
Activation function for deltas
@private
 */
func f(out float64) float64 {
	return (1 - out) * out
}
