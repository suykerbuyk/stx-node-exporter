package main

import (
	"fmt"

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
	for index, element := range enc.Enclosures {
		fmt.Println(element.Attributes.ID)
		fmt.Println(element.Attributes.Model)
		fmt.Println(element.Attributes.SasAddress)
		fmt.Println(index)
	}

}
