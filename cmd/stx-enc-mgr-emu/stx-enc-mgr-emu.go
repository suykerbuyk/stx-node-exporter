package main

import (
	_ "bufio"
	"io/ioutil"
	"net/http"
	_ "os"
)

var err error
var buf []byte

func sendJSON(w http.ResponseWriter, r *http.Request) {
	w.Write(buf)
}

func loadMessage(path string) error {
	buf = make([]byte, 65536)
	buf, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return err
}

func main() {
	err = loadMessage("../../api/stx-enc-mgr-metric.json")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/metric", sendJSON)
	if err = http.ListenAndServe(":9118", nil); err != nil {
		panic(err)
	}
}
