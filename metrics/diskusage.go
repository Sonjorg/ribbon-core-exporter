package metrics

import (
	"encoding/xml"
	"strconv"
	"strings"

	"core-exporter/lib"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	diskUsageName      = "hardDiskUsage"
	diskUsageUrlSuffix = "/restconf/data/sonusSystem:system/hardDiskUsage"
)

var DiskUsageMetric = lib.SonusMetric{
	Name:       diskUsageName,
	Processor:  processDiskUsage,
	URLGetter:  getDiskUsageUrl,
	APIMetrics: diskUsageMetrics,
	Repetition: lib.RepeatNone,
}

func getDiskUsageUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + diskUsageUrlSuffix
}

var diskUsageMetrics = map[string]*prometheus.Desc{
	"free_DiskSpace": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "disk", "free"),
		"Indicates free hard disk space (KBytes)",
		[]string{"server", "Partition"}, nil,
	),
	"used_DiskSpace": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "disk", "used"),
		"Indicates used hard disk (%)",
		[]string{"server", "Partition"}, nil,
	),
}

func processDiskUsage(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors        []*error
		disks = new(hardDiskUsageCollection)
	)

	err := xml.Unmarshal(*xmlBody, &disks)

	if err != nil {
		log.Errorf("Failed to deserialize DiskUsage XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: diskUsageName, Success: false, Errors: errors}
		return
	}

	for _, disk := range disks.HardDiskUsage {
		var freedisk, err = hardDiskUsage.freeToKbytes(*disk)
		if err != nil {
			log.Errorf("Failed to convert disk size (%q) to KB: %v", disk.FreeDiskSpace, err)
			errors = append(errors, &err)
			break
		}

		var useddisk, err2 = hardDiskUsage.usedToPrecent(*disk)
		if err2 != nil {
						log.Errorf("Failed to convert used disk size (%q) to percent, %v", disk.UsedDiskSpace, err2.Error())

			errors = append(errors, &err2)
			break
		}

		ctx.MetricChannel <- prometheus.MustNewConstMetric(diskUsageMetrics["free_DiskSpace"], prometheus.GaugeValue, freedisk, disk.ServerName, disk.Partition)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(diskUsageMetrics["used_DiskSpace"], prometheus.GaugeValue, useddisk, disk.ServerName, disk.Partition)

	}

	log.Info("hardDiskusage Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: diskUsageName, Success: true}
}

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <hardDiskUsage xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <serverName>NOGJHDO-SBC-01ta</serverName>
    <partition>/</partition>
    <totalDiskSpace>10218772 KBytes</totalDiskSpace>
    <freeDiskSpace>4387372 KBytes</freeDiskSpace>
    <usedDiskSpace>57%</usedDiskSpace>
    <role>primary</role>
    <syncStatus>unprotected</syncStatus>
    <syncCompletion>n/a</syncCompletion>
  </hardDiskUsage>
...
</collection>
*/

type hardDiskUsageCollection struct {
	HardDiskUsage []*hardDiskUsage `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 hardDiskUsage,omitempty"`
}

type hardDiskUsage struct {
	ServerName     string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 serverName"`
	Partition 	   string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 partition"`
	TotalDiskSpace string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 totalDiskSpace"`
	FreeDiskSpace  string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 freeDiskSpace"`
	UsedDiskSpace  string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 usedDiskSpace"`
	}

func (h hardDiskUsage) freeToKbytes() (float64, error) {
	var KBytes = strings.TrimSuffix(h.FreeDiskSpace, " KBytes")
	return strconv.ParseFloat(KBytes, 64)
}

func (h hardDiskUsage) usedToPrecent() (float64, error) {
	var KBytes = strings.TrimSuffix(h.UsedDiskSpace, "%")
	return strconv.ParseFloat(KBytes, 64)
}

func (h hardDiskUsage) totalToKbyes() (float64, error) {
	var gb = strings.TrimSuffix(h.TotalDiskSpace, " KBytes")
	return strconv.ParseFloat(gb, 64)
}
