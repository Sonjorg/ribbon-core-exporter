package metrics

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

var MgmtPortMetrics = map[string]*prometheus.Desc{
	"RxActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "rxActualBandwidth"),
		"Actual Rx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"TxActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "txActualBandwidth"),
		"Actual Tx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"RxPackets": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "rxPackets"),
		"The number of received packets.",
		[]string{"server", "port"}, nil,
	),
	"TxPackets": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "txPackets"),
		"The number of transmitted packets.",
		[]string{"server", "port"}, nil,
	),
  	"Linkstate": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "linkstate"),
		"0:null,1:Dsabld,2:PortDown,3:PortUp,4:DisabldNoLicense,5:EnbldPortDownInvalidSfpWrongSpeed,6:EnbldPortDownInvalidSfpNonSonus",
		[]string{"server", "port"}, nil,
	),
	"NegotiatedSpeed": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "negotiatedSpeed"),
		"1:speed10Mbps,2:speed100Mbps,3:speed1000Mbps,4:unknown,5:speed10000Mbps",
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

	for _, mgmtport := range mgmtPortStatus.MgmtPortStatuses  {
		ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["RxActualBandwidth"], prometheus.GaugeValue, mgmtport.RxActualBandwidth, mgmtport.CeName, mgmtport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["TxActualBandwidth"], prometheus.GaugeValue, mgmtport.TxActualBandwidth, mgmtport.CeName, mgmtport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["RxPackets"], prometheus.GaugeValue, mgmtport.RxPackets, mgmtport.CeName, mgmtport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["TxPackets"], prometheus.GaugeValue, mgmtport.TxPackets, mgmtport.CeName, mgmtport.PortName)
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["Linkstate"], prometheus.GaugeValue, convertLinkStateToNum(mgmtport.LinkState), mgmtport.CeName, mgmtport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(MgmtPortMetrics["NegotiatedSpeed"], prometheus.GaugeValue, convertSpeedToNum(mgmtport.NegotiatedSpeed), mgmtport.CeName, mgmtport.PortName)
	}

	log.Info("MgmtPort Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: MgmtPortName, Success: true}
}

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <mgmtPortStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:ORCA_SYSTEM="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01ta</ceName>
    <portName>mgt0</portName>
    <ifIndex>1</ifIndex>
    <macAddress>0:10:6b:5:59:de</macAddress>
    <negotiatedSpeed>speed1000Mbps</negotiatedSpeed>
    <linkState>admnEnabledPortUp</linkState>
    <duplexMode>full</duplexMode>
    <rxActualBandwidth>0</rxActualBandwidth>
    <txActualBandwidth>0</txActualBandwidth>
    <rxPackets>9481590</rxPackets>
    <txPackets>1333039</txPackets>
    <rxErrors>0</rxErrors>
    <txErrors>0</txErrors>
    <rxDropped>0</rxDropped>
    <txDropped>0</txDropped>
  </mgmtPortStatus>
*/


type MgmtPortStatusCollection struct {
	MgmtPortStatuses []*MgmtPortStatus `xml:"mgmtPortStatus"`
}

type MgmtPortStatus struct {
	CeName            string  `xml:"ceName"` 
	PortName          string  `xml:"portName"`
	LinkState         string  `xml:"linkState"`
	NegotiatedSpeed   string  `xml:"negotiatedSpeed"`
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

func convertSpeedToNum(negotiatedSpeed string) float64 {
	switch negotiatedSpeed {
	case "speed10Mbps":
		return 1
	case "speed100Mbps":
		return 2
	case "speed1000Mbps":
		return 3
	case "unknown":
		return 4
	case "speed10000Mbps":
		return 5

	}
	return 4
}
