package layer

import "gonn/synapse"

/**
Layer
 */
type Layer struct {
	Index int
	Nodes int
	Synapses []synapse.Synapse
}
