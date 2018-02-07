package main_interface

import (
	"gonn/net"
)

/**
Main Interface
 */
func NewNetwork(title string, moment, speed float64, bias bool, nodes []int) net.Network {
	var n = net.Network{}
	n.Create(title, moment, speed, bias, nodes)

	return n
}

func ReadNetwork(title string) net.Network {
	var n = net.Network{Title:title}
	n.Read()

	return n
}

/**
Link to default function GetResult
@public
 */
func GetResult(n net.Network, input []float64) []float64 {
	return n.GetResult(input);
}

/**
Link to default function Training
@public
 */
func Training(n net.Network, input, ideal [][]float64)  {
	n.Training(input, ideal)
}
