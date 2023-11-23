
package metrics

import (
	"encoding/xml"

	"core-exporter/lib"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	IpPolicingName      = "IpPolicingAlarmStatus"
	IpPolicingUrlSuffix = "/restconf/data/sonusOrca:oam/alarms/ipPolicingAlarmStatus"
)

var IpPolicingMetric = lib.SonusMetric{
	Name:       IpPolicingName,
	Processor:  processIpPolicing,
	URLGetter:  getIpPolicingUrl,
	APIMetrics: IpPolicingMetrics,
	Repetition: lib.RepeatNone,
}

func getIpPolicingUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + IpPolicingUrlSuffix
}

var IpPolicingMetrics = map[string]*prometheus.Desc{
	"AlarmDuration": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ipPolicing", "alarmDuration"),
		"Number of seconds the system {type} policer alarm has been at this level.",
		[]string{"systemName","alarmLevel", "type"}, nil,
	),
	"DiscardRate": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ipPolicing", "discardRate"),
		"Current rate of {type} discards for the system.",
		[]string{"systemName","alarmLevel", "type"}, nil,
	),
  "PacketsDiscarded": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ipPolicing", "packetsDiscarded"),
		"Total number of packets discarded by {type} policers on the system.",
		[]string{"systemName","alarmLevel", "type"}, nil,
	),
  "PacketsAccepted": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "ipPolicing", "packetsAccepted"),
		"Total number of {type} packets accepted on the system.",
		[]string{"systemName", "alarmLevel","type"}, nil,
	),
}

func processIpPolicing(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors        []*error
		IpPolicing = new(IpPolicingCollection)
	)

	err := xml.Unmarshal(*xmlBody, &IpPolicing)

	if err != nil {
		log.Errorf("Failed to deserialize IpPolicingAlarmStatus XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: IpPolicingName, Success: false, Errors: errors}
		return
	}

	for _, status := range IpPolicing.IpPolicingAlarms {
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.BadEthernetIpHeaderAlarmDuration, status.ServerName, status.BadEthernetIpHeaderAlarmLevel,"BadEthernetIpHeader")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.BadEthernetIpHeaderDiscardRate, status.ServerName, status.BadEthernetIpHeaderAlarmLevel,"BadEthernetIpHeader")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.BadEthernetIpHeaderPacketsDiscarded, status.ServerName, status.BadEthernetIpHeaderAlarmLevel,"BadEthernetIpHeader")
   		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.ArpAlarmDuration, status.ServerName, status.ArpAlarmLevel,"Arp")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.ArpDiscardRate, status.ServerName, status.ArpAlarmLevel,"Arp")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.ArpPacketsDiscarded, status.ServerName, status.ArpAlarmLevel,"Arp")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsAccepted"], prometheus.GaugeValue, status.ArpPacketsAccepted, status.ServerName, status.ArpAlarmLevel,"Arp")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.UFlowAlarmDuration, status.ServerName, status.UFlowAlarmLevel,"UFlow")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.UFlowDiscardRate, status.ServerName, status.UFlowAlarmLevel,"UFlow")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.UFlowPacketsDiscarded, status.ServerName, status.UFlowAlarmLevel,"UFlow")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsAccepted"], prometheus.GaugeValue, status.UFlowPacketsAccepted, status.ServerName, status.UFlowAlarmLevel,"UFlow")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.AclAlarmDuration, status.ServerName, status.AclAlarmLevel,"Acl")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.AclDiscardRate, status.ServerName, status.AclAlarmLevel,"Acl")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.AclPacketsDiscarded, status.ServerName, status.AclAlarmLevel,"Acl")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsAccepted"], prometheus.GaugeValue, status.AclPacketsAccepted, status.ServerName, status.AclAlarmLevel,"Acl")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.AggregateAlarmDuration, status.ServerName, status.AggregateAlarmLevel,"Aggregate")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.AggregateDiscardRate, status.ServerName, status.AggregateAlarmLevel,"Aggregate")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.AggregatePacketsDiscarded, status.ServerName, status.AggregateAlarmLevel,"Aggregate")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsAccepted"], prometheus.GaugeValue, status.AggregatePacketsAccepted, status.ServerName, status.AggregateAlarmLevel,"Aggregate")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.IpSecDecryptAlarmDuration, status.ServerName, status.IpSecDecryptAlarmLevel,"IpSecDecrypt")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.IpSecDecryptDiscardRate, status.ServerName, status.IpSecDecryptAlarmLevel,"IpSecDecrypt")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.IpSecDecryptPacketsDiscarded, status.ServerName, status.IpSecDecryptAlarmLevel,"IpSecDecrypt")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.MediaAlarmDuration, status.ServerName, status.MediaAlarmLevel,"Media")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.MediaDiscardRate, status.ServerName, status.MediaAlarmLevel,"Media")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.MediaPacketsDiscarded, status.ServerName, status.MediaAlarmLevel,"Media")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.RogueMediaAlarmDuration, status.ServerName, status.RogueMediaAlarmLevel,"RogueMedia")
		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.RogueMediaDiscardRate, status.ServerName, status.RogueMediaAlarmLevel,"RogueMedia")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.RogueMediaPacketsDiscarded, status.ServerName, status.RogueMediaAlarmLevel,"RogueMedia")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.DiscardRuleAlarmDuration, status.ServerName, status.DiscardRuleAlarmLevel,"DiscardRule")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.DiscardRuleDiscardRate, status.ServerName, status.DiscardRuleAlarmLevel,"DiscardRule")
   		ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.DiscardRulePacketsDiscarded, status.ServerName, status.DiscardRuleAlarmLevel,"DiscardRule")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["AlarmDuration"], prometheus.GaugeValue, status.SrtpDecryptAlarmDuration, status.ServerName, status.SrtpDecryptAlarmLevel,"SrtpDecrypt")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["DiscardRate"], prometheus.GaugeValue, status.SrtpDecryptDiscardRate, status.ServerName, status.SrtpDecryptAlarmLevel,"SrtpDecrypt")
    	ctx.MetricChannel <- prometheus.MustNewConstMetric(IpPolicingMetrics["PacketsDiscarded"], prometheus.GaugeValue, status.SrtpDecryptPacketsDiscarded, status.ServerName, status.SrtpDecryptAlarmLevel,"SrtpDecrypt")
	}

	log.Info("IpPolicingAlarmStatus Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: IpPolicingName, Success: true}
}

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <ipPolicingAlarmStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ORCA/1.0"  xmlns:ORCA="http://sonusnet.com/ns/mibs/SONUS-ORCA/1.0">
    <systemName>NOGJHDO-SBC-01t</systemName>
    <badEthernetIpHeaderAlarmLevel>noAlarm</badEthernetIpHeaderAlarmLevel>
    <badEthernetIpHeaderAlarmDuration>13823753</badEthernetIpHeaderAlarmDuration>
    <badEthernetIpHeaderDiscardRate>7</badEthernetIpHeaderDiscardRate>
    <badEthernetIpHeaderPacketsDiscarded>79637416</badEthernetIpHeaderPacketsDiscarded>
    <arpAlarmLevel>noAlarm</arpAlarmLevel>
    <arpAlarmDuration>17354099</arpAlarmDuration>
    <arpDiscardRate>0</arpDiscardRate>
    <arpPacketsDiscarded>0</arpPacketsDiscarded>
    <arpPacketsAccepted>811164</arpPacketsAccepted>
    <uFlowAlarmLevel>noAlarm</uFlowAlarmLevel>
.......
*/

type IpPolicingCollection struct {
	IpPolicingAlarms []*IpPolicing `xml:" ipPolicingAlarmStatus,omitempty"`
}

type IpPolicing struct {
	ServerName    string `xml:"systemName"`
	BadEthernetIpHeaderAlarmLevel string `xml:"badEthernetIpHeaderAlarmLevel"`
	BadEthernetIpHeaderAlarmDuration       float64   `xml:"badEthernetIpHeaderAlarmDuration"`
	BadEthernetIpHeaderDiscardRate    float64   `xml:"badEthernetIpHeaderDiscardRate"`
	BadEthernetIpHeaderPacketsDiscarded  float64   `xml:"badEthernetIpHeaderPacketsDiscarded"`
  	ArpAlarmLevel string `xml:"arpAlarmLevel"`
	ArpAlarmDuration       float64   `xml:"arpAlarmDuration"`
	ArpDiscardRate    float64   `xml:"arpDiscardRate"`
	ArpPacketsDiscarded  float64   `xml:"arpPacketsDiscarded"`
 	 ArpPacketsAccepted float64 `xml:"arpPacketsAccepted"`
	UFlowAlarmLevel       string   `xml:"uFlowAlarmLevel"`
	UFlowAlarmDuration    float64   `xml:"uFlowAlarmDuration"`
	UFlowDiscardRate  float64   `xml:"uFlowDiscardRate"`
 	 UFlowPacketsDiscarded float64 `xml:"uFlowPacketsDiscarded"`
	UFlowPacketsAccepted       float64   `xml:"uFlowPacketsAccepted"`
	AclAlarmLevel    string   `xml:"aclAlarmLevel"`
	AclAlarmDuration  float64   `xml:"aclAlarmDuration"`
  	AclDiscardRate float64 `xml:"aclDiscardRate"`
	AclPacketsDiscarded       float64   `xml:"aclPacketsDiscarded"`
	AclPacketsAccepted    float64   `xml:"aclPacketsAccepted"`
	AggregateAlarmLevel  string   `xml:"aggregateAlarmLevel"`
  	AggregateAlarmDuration float64 `xml:"aggregateAlarmDuration"`
	AggregateDiscardRate       float64   `xml:"aggregateDiscardRate"`
	AggregatePacketsDiscarded    float64   `xml:"aggregatePacketsDiscarded"`
	AggregatePacketsAccepted  float64   `xml:"aggregatePacketsAccepted"`
	IpSecDecryptAlarmLevel string `xml:"ipSecDecryptAlarmLevel"`
	IpSecDecryptAlarmDuration       float64   `xml:"ipSecDecryptAlarmDuration"`
	IpSecDecryptDiscardRate    float64   `xml:"ipSecDecryptDiscardRate"`
	IpSecDecryptPacketsDiscarded  float64   `xml:"ipSecDecryptPacketsDiscarded"`
  	MediaAlarmLevel string `xml:"mediaAlarmLevel"`
	MediaAlarmDuration       float64   `xml:"mediaAlarmDuration"`
	MediaDiscardRate    float64   `xml:"mediaDiscardRate"`
	MediaPacketsDiscarded  float64   `xml:"mediaPacketsDiscarded"`
  	RogueMediaAlarmLevel string `xml:"rogueMediaAlarmLevel"`
	RogueMediaAlarmDuration       float64   `xml:"rogueMediaAlarmDuration"`
	RogueMediaDiscardRate    float64   `xml:"rogueMediaDiscardRate"`
	RogueMediaPacketsDiscarded  float64   `xml:"rogueMediaPacketsDiscarded"`
  	DiscardRuleAlarmLevel string `xml:"discardRuleAlarmLevel"`
	DiscardRuleAlarmDuration       float64   `xml:"discardRuleAlarmDuration"`
	DiscardRuleDiscardRate    float64   `xml:"discardRuleDiscardRate"`
	DiscardRulePacketsDiscarded  float64   `xml:"discardRulePacketsDiscarded"`
  	SrtpDecryptAlarmLevel string `xml:"srtpDecryptAlarmLevel"`
	SrtpDecryptAlarmDuration       float64   `xml:"srtpDecryptAlarmDuration"`
	SrtpDecryptDiscardRate    float64   `xml:"srtpDecryptDiscardRate"`
	SrtpDecryptPacketsDiscarded  float64   `xml:"srtpDecryptPacketsDiscarded"`
}


