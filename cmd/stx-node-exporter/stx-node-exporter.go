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
	//ada := enc.Enclosures[0].Elements
	//count := len(ada)
	//fmt.Println("ada: ", ada, "count: ", count)
	//for idx1 := range enc.Enclosures {
	//	for idx2, elems := range enc.Enclosures[idx1].Elements {
	//		//fmt.Println(reflect.TypeOf(enc.Enclosures[index]))
	//		//fmt.Println(idx1, idx2, reflect.TypeOf(enc.Enclosures[idx1].Elements[idx2]))
	//		ada := enc.Enclosures[idx1].elements[idx2].ArrayDevice.Device
	//		NumberOfElems := len(enc.Enclosures.ArrayDevices.Device[0])
	//		fmt.Println("NumberOfElems:", NumberOfElems)
	//		//for idx3 := range enc.Enclosures[idx1].Elements[idx2].ArrayDevices.Device {
	//		//	fmt.Println(idx1, idx2, idx3, enc.Enclosures[idx1].Elements[idx2].ArrayDevices.Device[idx3])
	//		//}
	//	}
	//}

}
