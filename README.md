# stx-node-exporter
Translates sgt_enc_mgr JSON to Prometheus Node Exporter metrics

# Current state
## Two packages for use and reuse
### pkg/vendor - to be used with the makefile to support versioning
Makefile still broken
### pkg/encmgr
Handles communication to and from the stx-enc-mgr daemon.
## Executables
### cmd/stx-enc-mgr-emu
Emulates the behavior of the real daemon running on real hardware.  Useful for test and development.
### cmd/stx-node-exporter
What should become the node exporter itself.  Still not able to tie in the prometheus stuff.

## Examples
* https://godoc.org/github.com/kandoo/beehive/Godeps/_workspace/src/github.com/prometheus/client_golang/prometheus
* https://github.com/1and1/dellhw_exporter
* https://github.com/galexrt/dellhw_exporter/blob/master/collector/chassis.go
* https://github.com/mindprince/nvidia_gpu_prometheus_exporter/blob/master/main.go
## docs
* https://prometheus.io/docs/instrumenting/writing_exporters/


