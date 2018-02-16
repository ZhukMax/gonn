package benchmarks

import (
	"testing"
	"gonn/main_interface"
	"gonn/tests"
	"os"
	"gonn"
)

func BenchmarkCreateNewNetwork(b *testing.B) {
	var a = tests.GetTestNetwork()

	for i := 0; i < b.N; i++ {
		main_interface.NewNetwork(a.Title, a.Moment, a.Speed, false, a.Nodes)
	}
}

func BenchmarkReadNetwork(b *testing.B)  {
	var a = tests.GetTestNetwork()

	for i := 0; i < b.N; i++ {
		main_interface.ReadNetwork(a.Title)
	}
}

func BenchmarkGetResult(b *testing.B) {
	var a = tests.GetTestNetwork()
	var n = main_interface.ReadNetwork(a.Title)

	for i := 0; i < b.N; i++ {
		main_interface.GetResult(n, []float64{0, 1})
	}
}

func BenchmarkTraining(b *testing.B) {
	var a = tests.GetTestNetwork()
	var n = main_interface.ReadNetwork(a.Title)
	var input = tests.GetTrainingData()
	var ideal  = tests.GetTrainingIdeal()

	for i := 0; i < b.N; i++ {
		main_interface.Training(n, input, ideal, 1)
	}

	os.Remove(gonn.DataPath + "/" + n.File)
}
