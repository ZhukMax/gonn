package tests

import (
	"testing"
	"gonn/main_interface"
	"os"
)

func TestReadNetwork(t *testing.T)  {
	var a = getTestNetwork()
	main_interface.NewNetwork(a.Title, a.Moment, a.Speed, false, a.Nodes)
	var n = main_interface.ReadNetwork(a.Title)

	if len(n.Title) == 0 {
		t.Error("Length of Title is null")
	}
	if n.Title != a.Title {
		t.Error("Title must be like \"", a.Title, "\"")
	}
	if len(n.File) == 0 {
		t.Error("Length of File path is null")
	}
	if n.Moment != a.Moment {
		t.Error("Moment must be", a.Moment)
	}
	if n.Speed != a.Speed {
		t.Error("Speed must be", a.Speed)
	}
	if len(n.Structure) != len(a.Nodes) {
		t.Error("Wrong Structure of Network, must be", len(a.Nodes), "layers, but it's", len(n.Structure))
	}

	os.Remove(n.File)
}
