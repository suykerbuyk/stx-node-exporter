package collector

import (
	"fmt"
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

func collectDeviceArrayValues() ([]deviceStateValue, error) {
	devValues := []deviceStateValue{}
	for encIdx := range Enclosures.Enclosures {
		enc := &Enclosures.Enclosures[encIdx]
		encID := strings.ReplaceAll(Namespace+"_"+enc.Attributes.Model+"_"+enc.Attributes.Serial, " ", "")
		encID = sanitizeMetricString(encID)
		for _, dev := range enc.Elements.ArrayDevices.Device {
			if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
				// Current Array devices do not have a supported global status.
				continue
			}
			if dev.Status == encmgr.EncStatusCodeNoAccessAllowed {
				// Probably mapped to other controller, nothing to see here.
				continue
			}
			devIdStr := fmt.Sprintf("%03d", dev.Number)
			component := strings.ReplaceAll(dev.TypeStr+"_"+devIdStr, " ", "")
			value := deviceStateValue{
				Name:   encID,
				Value:  float64(dev.Status),
				Labels: map[string]string{"component": component},
			}
			devValues = append(devValues, value)
		}
	}
	return devValues, nil
}

// Update Prometheus metrics
func (a *arrayDeviceCollector) Update(ch chan<- prometheus.Metric) error {
	values, err := collectDeviceArrayValues()
	if err != nil {
		return err
	}
	for _, value := range values {
		a.current = prometheus.NewDesc(
			value.Name,
			devStatusHelpString,
			nil,
			value.Labels,
		)
		ch <- prometheus.MustNewConstMetric(
			a.current,
			prometheus.GaugeValue,
			value.Value,
		)
	}
	return nil
}
