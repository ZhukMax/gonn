package synapse

import (
	"time"
	"math/rand"
)

/**
Create new Synapse
@public
 */
func MakeSynapse(indexIn, indexOut int) Synapse {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var synapse = Synapse{}
	synapse.IndexIn = indexIn
	synapse.IndexOut = indexOut
	synapse.Weight = float64(r.Intn(1000)) / 100

	return synapse
}
