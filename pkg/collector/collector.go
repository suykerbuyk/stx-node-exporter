package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
)

// Namespace is the OEM prefix for all reported metrics keys.
const Namespace = "StxEnc"

// Factories is the list of all Collector interface implementations
var Factories = make(map[string]func() (Collector, error))

// Enclosures - the data structure periodically retrieved from stx-mgr
var Enclosures *encmgr.StxEncMgrMetrics

// Collector is the interface implemented by each metric type.
type Collector interface {
	Update(ch chan<- prometheus.Metric) error
}

type deviceStateValue struct {
	Name   string
	Value  float64
	Labels map[string]string
}

const devStatusHelpString string = "Device Status: 0-unsupported, 1-OK, 2-Critical, 3-Noncritical, 4-Unrecoverable, 5-NotInstalled, 6-Unknown, 7-NotAvailable, 8-NoAccess"

func sanitizeMetricString(s string) string {
	ret := strings.ReplaceAll(s, " ", "_")
	ret = strings.ReplaceAll(s, "-", "_")
	return ret
}

// FetchEnclosures gets/updates the enclosure data structures
func FetchEnclosures() error {
	var encs encmgr.StxEncMgrMetrics
	err := encs.ReadFromNetwork("http://localhost:9118/v1/metric")
	if err != nil {
		return err
	}
	Enclosures = &encs
	return nil
}
