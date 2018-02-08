package net

import (
	"gonn/layer"
)

/**
Network
 */
type Network struct {
	Title string
	File string
	Epoch int
	Iteration int
	Moment float64
	Speed float64
	Bias bool
	Structure []layer.Layer
	Error float64
}
