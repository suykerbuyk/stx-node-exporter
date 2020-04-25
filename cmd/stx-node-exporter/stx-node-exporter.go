package main

import (
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
)

func main() {
	var enc encmgr.StxEncMgrMetrics
	err := enc.ReadFromNetwork("http://localhost:9118/metric")
	if err != nil {
		panic(err)
	}
	encmgr.PrintJSONReport(&enc)
	err = enc.WriteToJSONFile("echo.json")
	if err != nil {
		panic(err)
	}
}
