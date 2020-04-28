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
	// for idx1 := range enc.Enclosures {
	// 	for idx2 := range enc.Enclosures[idx1].Elements {
	// 		//fmt.Println(reflect.TypeOf(enc.Enclosures[index]))
	// 		//fmt.Println(idx1, idx2, reflect.TypeOf(enc.Enclosures[idx1].Elements[idx2]).String())
	// 		for idx3 := range enc.Enclosures[idx1].Elements[idx2].ArrayDevice.Device {
	// 			fmt.Println(idx1, idx2, idx3, enc.Enclosures[idx1].Elements[idx2].ArrayDevice.Device[idx3])
	// 		}
	// 	}
	// }

}
