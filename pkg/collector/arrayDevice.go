package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
)

type arrayDeviceCollector struct {
	current *prometheus.Desc
}

func init() {
	Factories["ArrayDevice"] = NewArrayDeviceCollector
}

// NewArrayDeviceCollector returns a new arrayDeviceCollector
func NewArrayDeviceCollector() (Collector, error) {
	return &arrayDeviceCollector{}, nil
}

// Update Prometheus metrics
func (a *arrayDeviceCollector) Update(ch chan<- prometheus.Metric) error {
	for encIdx := range Enclosures.Enclosures {
		enc := &Enclosures.Enclosures[encIdx]
		encID := strings.ReplaceAll(Namespace+"_"+enc.Attributes.Model+"_"+enc.Attributes.Serial, " ", "")
		for _, dev := range enc.Elements.ArrayDevices.Device {
			if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
				continue
			}
			if dev.Status == encmgr.EncStatusCodeNoAccessAllowed {
				continue
			}
			out := prometheus.BuildFQName(encID, dev.TypeStr, "Status")
			val := float64(dev.Status)
			a.current = prometheus.NewDesc(
				out,
				"ArrayDevice status 0,1,2,3,4,5",
				nil,
				"Status",
			)
			ch <- prometheus.MustNewConstMetric(
				a.current, prometheus.GaugeValue, val)
		}
	}
	return nil
}
