package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("../../api/stx-enc-mgr-metric.json")
	//jsonFile, err := os.Open("pet-output.nojq")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var enc StxEncMgrJSON

	err = json.Unmarshal([]byte(byteValue), &enc)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)
	WriteJSONReportToFile(&enc, "echo.json")
}
