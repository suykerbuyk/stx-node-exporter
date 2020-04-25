package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func checkFileIO() {
	enc := StxEncMgrMetrics{}
	inFileName := "/../../api/stx-enc-mgr-metric.json"
	outFileName := "/echo.json"
	var err error
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inputJSONFile, err := filepath.EvalSymlinks(wd + inFileName)
	if err != nil {
		panic(err)
	}
	outputJSONFile, err := filepath.EvalSymlinks(wd + outFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read: " + inputJSONFile)
	if err = StxEncMetricsFromFile(&enc, inputJSONFile); err != nil {
		panic(err)
	}
	if err = WriteJSONReportToFile(&enc, outputJSONFile); err != nil {
		panic(err)
	}
	fmt.Println("Wrote: " + outputJSONFile)
}

func main() {
	// Basic HTTP GET request
	resp, err := http.Get("http://localhost:9118/metric")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	var enc StxEncMgrMetrics
	err = json.Unmarshal(body, &enc)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	PrintJSONReport(&enc)
}
