package tests

import (
	"testing"
	"gonn/main_interface"
	"os"
	"gonn"
)

func TestTraining(t *testing.T)  {
	var a = GetTestNetwork()
	var n = main_interface.NewNetwork(a.Title, a.Moment, a.Speed, false, a.Nodes)
	var input = GetTrainingData()
	var ideal  = GetTrainingIdeal()

	// First train for create Error
	main_interface.Training(n, input, ideal, 1)

	for i := 0; i < 10; i++ {
		// Re-read Net from file
		n := main_interface.ReadNetwork(a.Title)
		tmpError := n.Error

		// Train Net with save in file
		main_interface.Training(n, input, ideal, 1)

		// Read Net to variable for get new Error of Net
		tmpNet := main_interface.ReadNetwork(a.Title)

		if tmpError <= tmpNet.Error {
			t.Error("Error of net must to decrease. Last error:", tmpError, "New error:", tmpNet.Error)
		}
	}

	os.Remove(gonn.DataPath + "/" + n.File)
}
