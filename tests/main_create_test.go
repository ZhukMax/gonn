package tests

import (
	"testing"
	"gonn/main_interface"
	"os"
)

func TestNewNetworkWithoutBias(t *testing.T) {
	var a = getTestNetwork()
	var n = main_interface.NewNetwork(a.Title, a.Moment, a.Speed, false, a.Nodes)

	if len(n.Title) == 0 {
		t.Error("Length of Title is null")
	}
	if n.Title != a.Title {
		t.Error("Title must be like \"", a.Title, "\"")
	}
	if len(n.File) == 0 {
		t.Error("Length of File path is null")
	}
	if n.Epoch != 1 {
		t.Error("Epoch must be 1")
	}
	if n.Iteration != 1 {
		t.Error("Iteration must be 1")
	}
	if n.Moment != a.Moment {
		t.Error("Moment must be", a.Moment)
	}
	if n.Speed != a.Speed {
		t.Error("Speed must be", a.Speed)
	}
	if n.Bias {
		t.Error("Bias must be false")
	}
	if len(n.Structure) != len(a.Nodes) {
		t.Error("Wrong Structure of Network, must be", len(a.Nodes), "layers, but it's", len(n.Structure))
	}

	os.Remove(n.File)
}

func TestNewNetworkWithBias(t *testing.T) {
	var title = "Test"
	var moment = 0.01
	var speed = 0.05
	var nodes = []int{3, 3, 1}
	var n = main_interface.NewNetwork(title, moment, speed, true, nodes)

	if len(n.Title) == 0 {
		t.Error("Length of Title is null")
	}
	if n.Title != title {
		t.Error("Title must be like \"", title, "\"")
	}
	if len(n.File) == 0 {
		t.Error("Length of File path is null")
	}
	if n.Epoch != 1 {
		t.Error("Epoch must be 1")
	}
	if n.Iteration != 1 {
		t.Error("Iteration must be 1")
	}
	if n.Moment != moment {
		t.Error("Moment must be", moment)
	}
	if n.Speed != speed {
		t.Error("Speed must be", speed)
	}
	if !n.Bias {
		t.Error("Bias must be false")
	}
	if len(n.Structure) != len(nodes) {
		t.Error("Wrong Structure of Network, must be", len(nodes), "layers, but it's", len(n.Structure))
	}

	os.Remove(n.File)
}
