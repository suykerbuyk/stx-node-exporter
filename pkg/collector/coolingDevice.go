package collector

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
)

type coolingDeviceCollector struct {
	current *prometheus.Desc
}

type coolingDeviceStateValue struct {
	Name   string
	Value  float64
	Labels map[string]string
}

func init() {
	Factories["CoolingDevice"] = NewCoolingDeviceCollector
}

// NewCoolingDeviceCollector returns a new coolingDeviceCollector
func NewCoolingDeviceCollector() (Collector, error) {
	return &coolingDeviceCollector{}, nil
}

func collectCollingDeviceState() ([]coolingDeviceStateValue, error) {
	devValues := []coolingDeviceStateValue{}
	for encIdx := range Enclosures.Enclosures {
		enc := &Enclosures.Enclosures[encIdx]
		encID := strings.ReplaceAll(Namespace+"_"+enc.Attributes.Model+"_"+enc.Attributes.Serial, " ", "")
		encID = sanitizeMetricString(encID)
		for _, dev := range enc.Elements.CoolingDevices.Device {
			if dev.Status == encmgr.EncStatusCodeNoAccessAllowed {
				// Probably mapped to other controller, nothing to see here.
				continue
			}
			if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
				devIdStr := "GlobalStatus"
				component := strings.ReplaceAll(dev.TypeStr+"_"+devIdStr, " ", "")
				value := coolingDeviceStateValue{
					Name:   encID,
					Value:  float64(dev.GlobalStatus),
					Labels: map[string]string{"component": component},
				}
				devValues = append(devValues, value)
			} else {
				devIdStr := fmt.Sprintf("%03d", dev.Number)
				component := strings.ReplaceAll(dev.TypeStr+"_"+devIdStr, " ", "")
				value := coolingDeviceStateValue{
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
func (a *coolingDeviceCollector) Update(ch chan<- prometheus.Metric) error {
	values, err := collectCollingDeviceState()
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
