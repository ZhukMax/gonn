package net

import (
	"gonn/synapse"
	"fmt"
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
		// For the input layer
		if i == 0 {
			output = append(output, input)
		} else if i == len(n.Structure) - 1 {
			// If output layer
			break
		}

		var tmp []float64
		for node := 0; node < layer.Nodes; node++ {
			tmp = append(tmp, synapse.Activation(output[i], node, layer.Synapses))
		}

		output = append(output, tmp)
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
	var outputs = n.conclude(input)
	fmt.Print(outputs)
	var layers = len(outputs)

	return outputs[layers - 1]
}
