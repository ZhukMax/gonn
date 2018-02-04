package synapse

import "gonn"

/**
@public
 */
func SearchSynapse(i, o int, synapses []Synapse) float64 {
	var weight float64
	for y := range synapses {
		if synapses[y].IndexIn == i && synapses[y].IndexOut == o {
			weight = synapses[y].Weight
			break
		}
	}

	return weight
}

/**
@public
 */
func Activation(input []float64, node int, synapses []Synapse) float64 {
	var output float64
	for i, value := range input {
		output += value * SearchSynapse(i, node, synapses)
	}

	return gonn.Sigmoid(output)
}
