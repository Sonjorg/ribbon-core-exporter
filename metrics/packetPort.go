package metrics


import (
	"encoding/xml"

	"core-exporter/lib"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	packetPortName      = "PacketPortStatus"
	packetPortUrlSuffix = "/restconf/data/sonusSystem:system/sonusOrcaSystem:ethernetPort/packetPortStatus"
)

var PacketPortMetric = lib.SonusMetric{
	Name:       packetPortName,
	Processor:  processPacketPort,
	URLGetter:  getPacketPortUrl,
	APIMetrics: packetPortMetrics,
	Repetition: lib.RepeatNone,
}

func getPacketPortUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + packetPortUrlSuffix
}

var packetPortMetrics = map[string]*prometheus.Desc{
	"RxActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "RxActualBandwidth"),
		"Actual Rx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"TxActualBandwidth": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "txActualBandwidth"),
		"Actual Tx bandwidth in use on this port, bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"AvgRxActualBW": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "avgRxActualBW"),
		"Average Incoming Bandwidth used during KPI bin interval in bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"AvgTxActualBW": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "avgTxActualBW"),
		"Average Outgoing Bandwidth used during KPI bin interval in bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"PeakRxActualBW": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "peakRxActualBW"),
		"Peak Incoming Bandwidth used during KPI bin interval in bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"PeakTxActualBW": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "peakTxActualBW"),
		"Peak Outgoing Bandwidth used during KPI bin interval in bytes/sec.",
		[]string{"server", "port"}, nil,
	),
	"LinkState": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport_packetPort", "linkstate"),
		"0:null,1:Dsabld,2:PortDown,3:PortUp,4:DisabldNoLicense,5:EnbldPortDownInvalidSfpWrongSpeed,6:EnbldPortDownInvalidSfpNonSonus",
		[]string{"server", "port"}, nil,
	),
	"NegotiatedSpeed": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ethernetport", "negotiatedSpeed"),
		"1:speed10Mbps,2:speed100Mbps,3:speed1000Mbps,4:unknown,5:speed10000Mbps",
		[]string{"server", "port"}, nil,
	),
}

func processPacketPort(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors      []*error
		packetPorts = new(packetPortStatusCollection)
	)

	err := xml.Unmarshal(*xmlBody, &packetPorts)

	if err != nil {
		log.Errorf("Failed to deserialize packetPortStatus XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: packetPortName, Success: false, Errors: errors}
		return
	}

	for _, packetport := range packetPorts.PacketPortStatuses /*powerSupplies.PowerSupplyStatus*/ {

		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["RxActualBandwidth"], prometheus.GaugeValue, packetport.RxActualBandwidth, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["TxActualBandwidth"], prometheus.GaugeValue, packetport.TxActualBandwidth, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["AvgRxActualBW"], prometheus.GaugeValue, packetport.AvgRxActualBW, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["AvgTxActualBW"], prometheus.GaugeValue, packetport.AvgTxActualBW, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["PeakRxActualBW"], prometheus.GaugeValue, packetport.PeakRxActualBW, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["PeakTxActualBW"], prometheus.GaugeValue, packetport.PeakTxActualBW, packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["LinkState"], prometheus.GaugeValue, convertLinkStateToNum(packetport.LinkState), packetport.CeName, packetport.PortName)
		ctx.MetricChannel <- prometheus.MustNewConstMetric(packetPortMetrics["NegotiatedSpeed"], prometheus.GaugeValue, convertSpeedToNum(packetport.NegotiatedSpeed), packetport.CeName, packetport.PortName)

	}

	log.Info("PacketPortStatus Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: packetPortName, Success: true}
}

type packetPortStatusCollection struct {
	PacketPortStatuses []*PacketPortStatus `xml:"packetPortStatus,omitempty"`
}

type PacketPortStatus struct {
	CeName            string  `xml:"ceName"` //http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 ceName
	RxActualBandwidth float64 `xml:"rxActualBandwidth"`
	TxActualBandwidth float64 `xml:"txActualBandwidth"`
	LinkState         string  `xml:"linkState"`
	PortName          string  `xml:"portName"`
	AvgRxActualBW     float64 `xml:"avgRxActualBW"`
	AvgTxActualBW     float64 `xml:"avgTxActualBW"`
	PeakRxActualBW    float64 `xml:"peakRxActualBW"`
	PeakTxActualBW    float64 `xml:"peakTxActualBW"`
	NegotiatedSpeed   string  `xml:"negotiatedSpeed"`
}
