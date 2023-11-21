package main

import (
	"net/http"

	"core-exporter/config"
	"core-exporter/exporter"
	"core-exporter/lib"
	"core-exporter/metrics"

	"github.com/fatih/structs"
	"github.com/infinityworks/go-common/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	log            *logrus.Logger
	applicationCfg config.Config

	metricList = []lib.SonusMetric{
		metrics.DSPMetric,
		metrics.FanMetric,
		metrics.DiskStatusMetric,
		metrics.DiskUsageMetric,
	        metrics.MemoryMetric,
		metrics.CpuMetric,
		metrics.IPInterfaceMetric,
		metrics.PowerSupplyMetric,
		metrics.SipStatisticMetric,
		metrics.SipArsMetric,
		metrics.TGMetric,
		metrics.MgmtPortMetric,
		metrics.PacketPortMetric,
		metrics.SoftwareUpgradeMetric,
		metrics.CallCountMetric,
		metrics.SyncStatusMetric,
		metrics.IpPolicingMetric,
	}
)

func init() {
	applicationCfg = config.Init()
	log = logger.Start(&applicationCfg)
}

func main() {

	log.WithFields(structs.Map(applicationCfg)).Info("Starting Exporter")

	ex := exporter.Exporter{
		Metrics: metricList,
		Config:  applicationCfg,
	}

	// Register Metrics from each of the endpoints
	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(&ex)

	// Setup HTTP handler
	http.Handle(applicationCfg.MetricsPath(), promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		                <head><title>Ribbon SBC Exporter</title></head>
		                <body>
		                   <h1>Ribbon Metrics Exporter</h1>
						   <p>For more information, visit <a href=ssss>GitHub</a></p>
		                   <p><a href='` + applicationCfg.MetricsPath() + `'>Metrics</a></p>
		                   </body>
		                </html>
		              `))
	})
	log.Fatal(http.ListenAndServe(":"+applicationCfg.ListenPort(), nil))
}
