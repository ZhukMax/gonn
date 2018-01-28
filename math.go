package gonn

import (
	"math"
	"gonn/synapse"
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
func Delta(node int, out float64, synapses []synapse.Synapse, deltas []float64) float64 {
	if len(synapses) != len(deltas) {
		panic("Delta func need equally synapses and deltas")
	}

	var amount float64
	for i, delta := range deltas {
		amount += delta * synapse.SearchSynapse(node, i, synapses)
	}

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
