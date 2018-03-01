package net

import (
	"gonn/synapse"
)

/**
Conclude output data from all nodes
@private
 */
func (n Network) conclude(input []float64) [][]float64 {
	// Var for output from function
	var output [][]float64

	// Check input data
	if !n.checkInputData(input) {
		panic("The count of input nodes and input data is NOT equally")
	}

	// Conclude output data
	for i, layer := range n.Structure {
		if i == 0 {
			// For the input layer
			output = append(output, input)
		} else {
			var tmp []float64
			for node := 0; node < layer.Nodes; node++ {
				tmp = append(tmp, synapse.Activation(output[i - 1], node, n.Structure[i-1].Synapses))
			}

			output = append(output, tmp)
		}
	}

	return output
}

/**
Check that the count of input nodes and input data is equally
@private
 */
func (n Network) checkInputData(input []float64) bool {
	var countInputNodes = n.Structure[0].Nodes
	var countInputData  = len(input)
	if n.Bias {
		countInputData++
	}
	if countInputData != countInputNodes {
		return false
	}

	return true
}

func (n Network) GetResult(input []float64) []float64 {
	var output = n.conclude(input)
	var layers = len(output)

	return output[layers - 1]
}
