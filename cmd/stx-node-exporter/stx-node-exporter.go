package main

import (
	"flag"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
	"github.com/suykerbuyk/stx-node-exporter/pkg/version"
)

const (
	namespace = "stx_enc"
)

var (
	endPoint = flag.String("metrics_port", ":9110", "Port to listen on for metric request")
)

// Collector - holds and stages metrics for export
type Collector struct {
	fanSpeed prometheus.Gauge
}

//NewCollector - builds a new collector
func NewCollector() *Collector {
	return &Collector{
		fanSpeed: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "fan_speed",
				Help:      "Real Fan Speed",
			},
		),
	}
}

// Describe - sets the meta data descriptions for each exported metric
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.fanSpeed.Desc()
}
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.fanSpeed.Set(float64(10))
}
func main() {
	var enclosures encmgr.StxEncMgrMetrics
	fmt.Println(version.Print("MyTest"))
	err := enclosures.ReadFromNetwork("http://localhost:9118/v1/metric")
	if err != nil {
		panic(err)
	}
	//encmgr.PrintJSONReport(&enc)
	err = enclosures.WriteToJSONFile("echo.json")
	if err != nil {
		panic(err)
	}

	for encIdx := range enclosures.Enclosures {
		enc := &enclosures.Enclosures[encIdx]
		nodeID := enc.Attributes.Model + "_" + enc.Attributes.Serial
		fmt.Println(nodeID)
		for _, dev := range enc.Elements.ArrayDevices.Device {
			if dev.Status != encmgr.EncStatusCodeNoAccessAllowed {
				if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
					fmt.Println(nodeID, dev.TypeStr, "GlobalStatus: ", dev.GlobalStatus, dev.GlobalStatusStr)
				} else {
					fmt.Println(nodeID, dev.TypeStr, dev.Number, "Status: ", dev.Status, "=", dev.StatusStr)
				}

			}
		}
		for _, dev := range enc.Elements.PowerSupplies.Device {
			if dev.Status != encmgr.EncStatusCodeNoAccessAllowed {
				if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
					fmt.Println(nodeID, dev.TypeStr, "GlobalStatus: ", dev.GlobalStatus, dev.GlobalStatusStr)
				} else {
					fmt.Println(nodeID, dev.TypeStr, dev.Number, "Status: ", dev.Status, "=", dev.StatusStr)
				}

			}
		}
		for _, dev := range enc.Elements.CoolingDevices.Device {
			if dev.Status != encmgr.EncStatusCodeNoAccessAllowed {
				if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
					fmt.Println(nodeID, dev.TypeStr, "GlobalStatus: ", dev.GlobalStatus, dev.GlobalStatusStr)
				} else {
					fmt.Println(nodeID, dev.TypeStr, dev.Number, "Status: ", dev.Status, "=", dev.StatusStr)
					fmt.Println(nodeID, dev.TypeStr, dev.Number, "Fan Speed:", dev.ActualSpeed, "Failure:", dev.Failure)
				}

			}
		}
	}
}
