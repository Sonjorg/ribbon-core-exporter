package metrics

//memory metrics
//The memory utilization for the current interval period.

import (
	"encoding/xml"

	"core-exporter/lib"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	MgmtPortName      = "MgmtPortStatus"
	MgmtPortUrlSuffix = "/restconf/data/sonusSystem:system/sonusOrcaSystem:ethernetPort/mgmtPortStatus"
)

var MgmtPortMetric = lib.SonusMetric{
	Name:       MgmtPortName,
	Processor:  processMgmtPort,
	URLGetter:  getMgmtPortUrl,
	APIMetrics: MgmtPortMetrics,
	Repetition: lib.RepeatNone,
}

func getMgmtPortUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + MgmtPortUrlSuffix
}

/*
<averageSwap>0</averageSwap>

	    <highSwap>0</highSwap>
	    <lowSwap>0</lowSwap>

	    	rxActualBandwidth	Actual Rx bandwidth in use on this port, bytes/sec.
		txActualBandwidth	Actual Tx bandwidth in use on this port, bytes/sec.
		rxPackets	The number of received packets.
		txPackets	The number of transmitted packets.
		linkState<
*/
var MgmtPortMetrics = map[string]*prometheus.Desc{
	"rxActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "rxActualBandwidth"),
		"Actual Rx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"txActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "txActualBandwidth"),
		"Actual Tx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"rxPackets": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "rxPackets"),
		"The number of received packets.",
		[]string{"server", "port"}, nil,
	),
	"txPackets": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "txPackets"),
		"The number of transmitted packets.",
		[]string{"server", "port"}, nil,
	),
  "linkstate": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "linkstate"),
		"0:null,1:Dsabld,2:PortDown,3:PortUp,4:DisabldNoLicense,5:EnbldPortDownInvalidSfpWrongSpeed,6:EnbldPortDownInvalidSfpNonSonus",
		[]string{"server", "port"}, nil,
	),
}

func processMgmtPort(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors         []*error
		mgmtPortStatus = new(MgmtPortStatusCollection)
	)

	err := xml.Unmarshal(*xmlBody, &mgmtPortStatus)

	if err != nil {
		log.Errorf("Failed to deserialize MgmtPort XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: MgmtPortName, Success: false, Errors: errors}
		return
	}

	for _, mgmtport := range mgmtPortStatus.MgmtPortStatuses /*powerSupplies.PowerSupplyStatus*/ {
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MgmtPortMetrics["rxActualBandwidth"], prometheus.GaugeValue, mgmtport.RxActualBandwidth, mgmtport.CeName, mgmtport.PortName))
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MgmtPortMetrics["txActualBandwidth"], prometheus.GaugeValue, mgmtport.TxActualBandwidth, mgmtport.CeName, mgmtport.PortName))
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MgmtPortMetrics["rxPackets"], prometheus.GaugeValue, mgmtport.RxPackets, mgmtport.CeName, mgmtport.PortName))
		ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MgmtPortMetrics["txPackets"], prometheus.GaugeValue, mgmtport.TxPackets, mgmtport.CeName, mgmtport.PortName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MgmtPortMetrics["linkstate"], prometheus.GaugeValue, convertLinkStateToNum(mgmtport.LinkState), mgmtport.CeName, mgmtport.PortName))
	}

	log.Info("MgmtPort Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: MgmtPortName, Success: true}
}

type MgmtPortStatusCollection struct {
	MgmtPortStatuses []*MgmtPortStatus `xml:"mgmtPortStatus"`
}

type MgmtPortStatus struct {
	CeName            string  `xml:"ceName"` //http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 ceName
	PortName          string  `xml:"portName"`
	LinkState         string  `xml:"linkState"`
	RxActualBandwidth float64 `xml:"rxActualBandwidth"`
	TxActualBandwidth float64 `xml:"txActualBandwidth"`
	RxPackets         float64 `xml:"rxPackets"`
	TxPackets         float64 `xml:"txPackets"`
}

func convertLinkStateToNum(linkstate string) float64 {
	switch linkstate {
	case "null":
    return 0;
	case "admnDisabled":
    return 1;
	case "admnEnabledPortDown":
    return 2;
	case "admnEnabledPortUp":
    return 3;
	case "admnDisabledNoLicense":
    return 4;
	case "admnEnabledPortDownInvalidSfpWrongSpeed":
    return 5;
	case "admnEnabledPortDownInvalidSfpNonSonus":
    return 6;
	}
return 0;
}
