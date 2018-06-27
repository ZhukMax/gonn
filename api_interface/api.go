package api_interface

import (
	"net/http"
	"fmt"
	"gonn/main_interface"
	"strconv"
	"gonn"
	"encoding/json"
)

func parseParamToFloat(param string) float64 {
	output, err := strconv.ParseFloat(param, 64)
	if err == nil {
		gonn.Check(err)
	}

	return output
}

func parseNodesParam(param string) []int {
	var nodes []int
	b := []byte(param)
	err := json.Unmarshal(b, &nodes)
	if err == nil {
		gonn.Check(err)
	}

	return nodes
}

/**
Web Interface
 */
func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	var params = r.URL.Query()
	
	var title = params.Get("title")
	var moment = parseParamToFloat(params.Get("moment"))
	var speed = parseParamToFloat(params.Get("speed"))

	bias, err := strconv.ParseBool(params.Get("bias"))
	if err == nil {
		gonn.Check(err)
	}

	var nodes = parseNodesParam(params.Get("nodes"))
	var network = main_interface.NewNetwork(title, moment, speed, bias, nodes)

	fmt.Fprint(w, network)
}

// TODO
func GetResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

// TODO
func Training(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

// todo:
func showNetwork(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

func Server(config gonn.Config)  {
	http.HandleFunc("/new", createHandler)
	http.ListenAndServe(":" + config.Port, nil)
}
