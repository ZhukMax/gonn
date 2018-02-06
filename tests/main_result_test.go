package tests

import (
	"gonn/main_interface"
	"testing"
	"os"
)

func TestGetResult(t *testing.T)  {
	var a = getTestNetwork()
	var n = main_interface.NewNetwork(a.Title, a.Moment, a.Speed, false, a.Nodes)
	var result = main_interface.GetResult(n, []float64{0, 1})

	if len(result) == 0 {
		t.Error("Result can't be empty array")
	}

	os.Remove(n.File)
}
