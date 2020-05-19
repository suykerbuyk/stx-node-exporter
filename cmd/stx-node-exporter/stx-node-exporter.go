package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/suykerbuyk/stx-node-exporter/pkg/collector"
	"github.com/suykerbuyk/stx-node-exporter/pkg/encmgr"
	"github.com/suykerbuyk/stx-node-exporter/pkg/flagutil"
)

// StxCollector contains the collectors to be used
type StxCollector struct {
	lastCollectTime time.Time
	collectors      map[string]collector.Collector
}

// Describe implements the prometheus.Collector interface.
func (s StxCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- scrapeDurationDesc
	ch <- scrapeSuccessDesc
}

// Collect implements the prometheus.Collector interface.
func (s StxCollector) Collect(ch chan<- prometheus.Metric) {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(s.collectors))
	for name, c := range s.collectors {
		go func(name string, c collector.Collector) {
			execute(name, c, ch)
			waitGroup.Done()
		}(name, c)
	}
	waitGroup.Wait()
}
func execute(name string, c collector.Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := c.Update(ch)
	duration := time.Since(begin)
	var success float64

	if err != nil {
		log.Errorf("%s collector failed after %fs: %s", name, duration.Seconds(), err)
		success = 0
	} else {
		log.Debugf("%s collector succeeded after %fs.", name, duration.Seconds())
		success = 1
	}
	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, duration.Seconds(), name)
	ch <- prometheus.MustNewConstMetric(scrapeSuccessDesc, prometheus.GaugeValue, success, name)
}
func loadCollectors(list string) (map[string]collector.Collector, error) {
	collectors := map[string]collector.Collector{}
	for _, name := range strings.Split(list, ",") {
		fn, ok := collector.Factories[name]
		if !ok {
			return nil, fmt.Errorf("collector '%s' not available", name)
		}
		c, err := fn()
		if err != nil {
			return nil, err
		}
		collectors[name] = c
	}
	return collectors, nil
}

var (
	scrapeDurationDesc = prometheus.NewDesc(
		//prometheus.BuildFQName(collector.Namespace, "scrape", "collector_duration_seconds"),
		"fred",
		"stx_node_exporter: Duration of a collector scrape.",
		//[]string{"collector"},
		[]string{"wilma"},
		//nil,
		map[string]string{"label1": "value1"},
	)
	scrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(collector.Namespace, "scrape", "collector_success"),
		"stx_node_exporter: Whether a collector succeeded.",
		[]string{"collector"},
		nil,
	)
)

const (
	defaultCollectors = "ArrayDevice"
)

//CmdLineOpts - runtime options
type CmdLineOpts struct {
	version           bool
	help              bool
	showCollectors    bool
	logLevel          string
	enabledCollectors string
	exportAddr        string
	exportPort        string
	encMgrAddr        string
	encMgrPort        string
}

var (
	log                 = logrus.New()
	opts                CmdLineOpts
	stxEncExporterFlags = flag.NewFlagSet("stx_node_exporter", flag.ExitOnError)
)

func init() {
	stxEncExporterFlags.BoolVar(&opts.help, "help", false, "Show help menu")
	stxEncExporterFlags.BoolVar(&opts.version, "version", false, "Show version information")
	stxEncExporterFlags.BoolVar(&opts.showCollectors, "collectors.show", false, "Only output the list of available collectors.")
	stxEncExporterFlags.StringVar(&opts.enabledCollectors, "collectors.enabled", defaultCollectors, "Comma separated list of collectors to enable")
	stxEncExporterFlags.StringVar(&opts.logLevel, "logLevel", "INFO", "Enable log output level (trace,debug,info, warn,error,fatal)")
	stxEncExporterFlags.StringVar(&opts.exportPort, "exportPort", "9110", "The port to serve metrics from")
	stxEncExporterFlags.StringVar(&opts.encMgrPort, "encMgrPort", "9118", "The port we query the stx-enc-mgr")

	// Define the usage function
	stxEncExporterFlags.Usage = usage

	if err := stxEncExporterFlags.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n", os.Args[0])
	stxEncExporterFlags.PrintDefaults()
	os.Exit(0)
}

func main() {
	if err := flagutil.SetFlagsFromEnv(stxEncExporterFlags, "STX_ENC_EXPORTER"); err != nil {
		log.Fatal(err)
	}
	if opts.help {
		usage()
	}
	if opts.version {
		fmt.Fprintln(os.Stdout, version.Print("stx_node_exporter 0.0"))
		os.Exit(0)
	}
	if opts.showCollectors {
		collectorNames := make(sort.StringSlice, 0, len(collector.Factories))
		for n := range collector.Factories {
			collectorNames = append(collectorNames, n)
		}
		collectorNames.Sort()
		fmt.Printf("Available collectors:\n")
		for _, n := range collectorNames {
			fmt.Printf(" - %s\n", n)
		}
		return
	}

	log.Out = os.Stdout
	level, err := logrus.ParseLevel(opts.logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)
	log.Infoln("Build context", version.BuildContext())
	err = collector.FetchEnclosures()
	if err != nil {
		log.Fatal(err)
	}
	collectors, err := loadCollectors(opts.enabledCollectors)
	if err != nil {
		log.Fatalf("Couldn't load collectors: %s", err)
	}
	log.Infof("Enabled collectors:")
	for n := range collectors {
		log.Infof("collector: %s", n)
	}
	if err = prometheus.Register(StxCollector{lastCollectTime: time.Now(), collectors: collectors}); err != nil {
		log.Fatalf("Couldn't register collector: %s", err)
	}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:      log,
			ErrorHandling: promhttp.ContinueOnError,
		})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Seagate Enclosure Exporter</title></head>
			<body>
			<h1>Seagate Exporter</h1>
			<p><a href="` + "/metrics" + `">Metrics</a></p>
			</body>
			</html>`))
	})

	if err := http.ListenAndServe(opts.exportPort, nil); err != nil {
		log.Fatal(err)
	}

}

func protoType() {
	err := collector.FetchEnclosures()
	if err != nil {
		panic(err)
	}
	//encmgr.PrintJSONReport(&enc)
	err = collector.Enclosures.WriteToJSONFile("echo.json")
	if err != nil {
		panic(err)
	}

	for encIdx := range collector.Enclosures.Enclosures {
		enc := &collector.Enclosures.Enclosures[encIdx]
		nodeID := collector.Namespace + "_" + enc.Attributes.Model + "_" + enc.Attributes.Serial
		var out string
		fmt.Println(nodeID)
		for _, dev := range enc.Elements.ArrayDevices.Device {
			if dev.Status != encmgr.EncStatusCodeNoAccessAllowed {
				if dev.Number == encmgr.EncDeviceTypeGlobalStatus {
					out = prometheus.BuildFQName(nodeID, dev.TypeStr, "GlobalStatus")
					fmt.Println(out, nodeID, dev.TypeStr, "GlobalStatus: ", dev.GlobalStatus, dev.GlobalStatusStr)
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
