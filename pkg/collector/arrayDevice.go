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
				// Current Array devices do not have a supported global status.
				continue
			}
			if dev.Status == encmgr.EncStatusCodeNoAccessAllowed {
				// Probably mapped to other controller, nothing to see here.
				continue
			}
			key := prometheus.BuildFQName(encID, dev.TypeStr, "Status")
			val := float64(dev.Status)
			lab1 := map[string]string{"label1": "value1"}
			a.current = prometheus.NewDesc(
				key,
				"ArrayDevice status 0,1,2,3,4,5",
				[]string{"Status", "status"},
				lab1,
			)
			ch <- prometheus.MustNewConstMetric(
				a.current, prometheus.GaugeValue, val)
		}
	}

	return nil
}
