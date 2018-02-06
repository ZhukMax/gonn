package tests

type TestNetwork struct {
	Title string
	Moment float64
	Speed float64
	Nodes []int
}

func getTestNetwork() TestNetwork {
	return TestNetwork{
		Title: "Test",
		Moment: 0.01,
		Speed: 0.05,
		Nodes: []int{2, 2, 1}}
}

func getTrainingData() [][]float64 {
	return [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1}}
}

func getTrainingIdeal() [][]float64 {
	return [][]float64{
		{0},
		{1},
		{1},
		{0}}
}
