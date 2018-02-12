package net

import (
	"math"
	"gonn"
	"gonn/synapse"
)

/**
@public
 */
func (n *Network) Training(input, ideal [][]float64)  {
	for i, data := range input {
		n.trainingIteration(data, ideal[i])
	}

	n.Epoch++
	n.Iteration = 1
}

/**
@private
 */
func (n *Network) trainingIteration(input, ideal []float64) {
	var result = n.conclude(input)
	var deltas = n.getDeltas(result, ideal)

	// Loop of Layers
	for i, l := range n.Structure {
		for _, s := range l.Synapses {
			// Gradient of synapse
			var gradient = deltas[i + 1][s.IndexOut] * result[i][s.IndexIn]

			// New difference
			s.Diff = gradient * n.Speed + n.Moment * s.Diff

			// Change synapse's weight
			s.Weight = s.Weight + s.Diff
		}
	}

	// Set new error of network
	n.setError(n.GetResult(input), ideal)
}

/**
Change Iteration of net
@private
 */
func (n *Network) setIteration() {
	n.Iteration++
}

func (n *Network) setError(results, purposes []float64) {
	var localError = n.Error * float64(n.Iteration)

	if len(results) != len(purposes) {
		panic("Wrong count of results and purposes for set a new NetError")
	}
	for i, result := range results {
		localError += math.Pow(purposes[i] - result, 2)
	}

	n.setIteration()

	// Change Error of net
	n.Error = localError / float64(n.Iteration)
}

/**
@private
@return [][]float64
 */
func (n Network) getDeltas(result [][]float64, ideal []float64) [][]float64 {
	var deltas [][]float64
	var outLayerNum = len(n.Structure) - 1
	for i := 0; i <= outLayerNum; i++ {
		deltas = append(deltas, []float64{})
	}

	for k := outLayerNum; k > 0; k-- {
		var localLayer = n.Structure[k]
		var localDeltas []float64

		for i := 0; i < localLayer.Nodes; i++ {
			if k == outLayerNum {
				// For output layer
				localDeltas = append(localDeltas, gonn.DeltaOut(result[k][i], ideal[i]))
			} else {
				// For hidden layers
				var amount float64
				for y, delta := range deltas[k + 1] {
					amount += delta * synapse.SearchSynapse(i, y, localLayer.Synapses)
				}
				localDeltas = append(localDeltas, gonn.Delta(result[k][i], amount))
			}
			deltas[k] = localDeltas
		}
	}

	return deltas
}
