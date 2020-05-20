package collector

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
)

type powerSupplyCollector struct {
	current *prometheus.Desc
}

func init() {
	Factories["PowerSupply"] = NewPowerSupplyCollector
}

// NewPowerSupplyCollector returns a new powerSupplyCollector
func NewPowerSupplyCollector() (Collector, error) {
	return &powerSupplyCollector{}, nil
}

func collectPowerSupplyValues() ([]deviceStateValue, error) {
	devValues := []deviceStateValue{}
	for encIdx := range Enclosures.Enclosures {
		enc := &Enclosures.Enclosures[encIdx]
		encID := strings.ReplaceAll(Namespace+"_"+enc.Attributes.Model+"_"+enc.Attributes.Serial, " ", "")
		encID = sanitizeMetricString(encID)
		for _, dev := range enc.Elements.PowerSupplies.Device {
			if dev.Status == encmgr.EncStatusCodeNoAccessAllowed {
				// Probably mapped to other controller, nothing to see here.
				continue
			}
			if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
				devIdStr := "GlobalStatus"
				component := strings.ReplaceAll(dev.TypeStr+"_"+devIdStr, " ", "")
				value := deviceStateValue{
					Name:   encID,
					Value:  float64(dev.GlobalStatus),
					Labels: map[string]string{"component": component},
				}
				devValues = append(devValues, value)
			} else {
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
	}
	return devValues, nil
}

// Update Prometheus metrics
func (a *powerSupplyCollector) Update(ch chan<- prometheus.Metric) error {
	values, err := collectPowerSupplyValues()
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
