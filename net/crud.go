package net

import (
	"gonn/layer"
	"encoding/json"
	"os"
	"io/ioutil"
	"gonn/synapse"
	"encoding/hex"
	"crypto/md5"
	"gonn"
)

/**
@private
 */
func (n *Network) setFile(title string) {
	hasher := md5.New()
	hasher.Write([]byte(title))
	n.File = hex.EncodeToString(hasher.Sum(nil)) + ".json"
}

/**
@private
 */
func (n *Network) makeStructure(nodes []int) {
	countLayers := len(nodes)

	for i, value := range nodes {
		// Don't add synapses for last layer
		if countLayers == i {
			break
		}

		if n.Bias {
			value++
		}

		l := layer.Layer{}
		l.Index = i
		l.Nodes = value

		// Create synapses for all nodes in layer
		for y := 0; y < value; y++ {
			for y2 := 0; y2 < nodes[i + 1]; y2++ {
				// Add to layer new synapse
				l.Synapses = append(l.Synapses, synapse.MakeSynapse(y, y2))
			}
		}

		// Add to Structure new layer
		n.Structure = append(n.Structure, l)
	}
}

func (n *Network) Create(title string, moment, speed float64, bias bool, nodes []int) {
	n.Title = title
	n.setFile(title)
	n.Epoch = 1
	n.Iteration = 1
	n.Moment = moment
	n.Speed = speed
	n.Bias = bias
	n.makeStructure(nodes)
	n.Save()
}

func (n *Network) Read() {
	n.setFile(n.Title)

	// read file
	data, err := ioutil.ReadFile(gonn.DataPath + "/" + n.File)
	gonn.Check(err)

	json.Unmarshal([]byte(data), &n)
}

func (n Network) Save() {
	f, err := os.Create(gonn.DataPath + "/" + n.File)
	gonn.Check(err)

	NetworkJson, _ := json.Marshal(n)
	f.Write([]byte(NetworkJson))
}
