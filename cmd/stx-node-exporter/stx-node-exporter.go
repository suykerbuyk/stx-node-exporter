package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func checkFileIO() {
	enc := StxEncMgrJSON{}
	inFileName := "/../../api/stx-enc-mgr-metric.json"
	outFileName := "/echo.json"
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inputJSONFile, err := filepath.EvalSymlinks(wd + inFileName)
	if err != nil {
		panic(err)
	}
	outputJSONFile, err := filepath.EvalSymlinks(wd + outFileName)
	fmt.Println("Read: " + inputJSONFile)
	if err = ReadJSONReportFromFile(&enc, inputJSONFile); err == nil {
		err = WriteJSONReportToFile(&enc, outputJSONFile)
	}
	if err != nil {
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

	fmt.Printf("%s\n", body)
}
