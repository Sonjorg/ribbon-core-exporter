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
	diskStatusName      = "DiskStatus"
	DiskStatusUrlSuffix = "/restconf/data/sonusSystem:system/hardDiskStatus"
)

var DiskStatusMetric = lib.SonusMetric{
	Name:       diskStatusName,
	Processor:  processDiskStatus,
	URLGetter:  getDiskStatusUrl,
	APIMetrics: diskStatusMetrics,
	Repetition: lib.RepeatNone,
	MetricArray: MetricArray,

}

func getDiskStatusUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + DiskStatusUrlSuffix
}

var diskStatusMetrics = map[string]*prometheus.Desc{
	"Disk_Status": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "disk", "status"),
		"Harddisk status that indicates if the disk is online/failed 1=online",
		[]string{"server", "productId"}, nil,
	),
	"Disk_healthtest": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "disk", "healthTest"),
		"Pass or Fail indicating the overall health of the device. 1=passed",
		[]string{"server"}, nil,
	),
	"Disk_Size": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "disk", "size"),
		"Capacity of the disk in the server",
		[]string{"server"}, nil,
	),
}

func processDiskStatus(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors        []*error
		disks = new(hardDiskStatusCollection)
	)

	err := xml.Unmarshal(*xmlBody, &disks)

	if err != nil {
		log.Errorf("Failed to deserialize DiskStatus XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: diskStatusName, Success: false, Errors: errors}
		return
	}

	for _, disk := range disks.HardDiskStatus {
		var diskSpace, err = hardDiskStatus.capacityToGB(*disk)
		if err != nil {
			log.Errorf("Failed to convert disk size (%q) to GB: %v", disk.Capacity, err)
			errors = append(errors, &err)
			break
		}

		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(diskStatusMetrics["Disk_Status"], prometheus.GaugeValue, disk.statusToFloat(), disk.ServerName, disk.ProductId))
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(diskStatusMetrics["Disk_healthtest"], prometheus.GaugeValue, disk.healtToFloat(), disk.ServerName))
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(diskStatusMetrics["Disk_Size"], prometheus.GaugeValue, diskSpace, disk.ServerName))
	}

	log.Info("Diskstatus Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: diskStatusName, Success: true}
}

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <hardDiskStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <serverName>NOGJHDO-SBC-01ta</serverName>
    <diskNumber>0</diskNumber>
    <present>true</present>
    <productId>ATA WDC PC SA530 SDA</productId>
    <revision>2100</revision>
    <capacity>476 GB</capacity>
    <diskStatus>online</diskStatus>
    <healthTest>PASSED</healthTest>
    <diskLifeRemaining>n/a</diskLifeRemaining>
  </hardDiskStatus>
...
</collection>
*/

type hardDiskStatusCollection struct {
	HardDiskStatus []*hardDiskStatus `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 hardDiskStatus,omitempty"`
}

type hardDiskStatus struct {
	ServerName    string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 serverName"`
	ProductId 	  string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 productId"`
	Present       bool   `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 present"`
	Capacity      string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 capacity"`
	DiskStatus    string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 diskStatus"`
	HealthTest    string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 healthTest"`
}

//convert disksize to GB value
func (h hardDiskStatus) capacityToGB() (float64, error) {
	var gb = strings.TrimSuffix(h.Capacity, " GB")
	return strconv.ParseFloat(gb, 64)
}

//disk status Online / offline
func (h hardDiskStatus) statusToFloat() float64 {
	if h.DiskStatus == "online" {
		return 1
	} else {
		return 0
	}
}
//disk healtcheck status
func (h hardDiskStatus) healtToFloat() float64 {
	if h.HealthTest == "PASSED" {
		return 1
	} else {
		return 0
	}
}