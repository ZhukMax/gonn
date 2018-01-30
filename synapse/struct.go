package synapse

/**
Synapse
 */
type Synapse struct {
	// indexes
	IndexIn int
	IndexOut int

	// weight of synapse
	Weight float64

	// Last difference synapse's weight
	Diff float64
}
