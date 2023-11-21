// +build test

package metrics


/*
//memory metrics
//The memory utilization for the current interval period.

import (
  "encoding/xml"

  "core-exporter/lib"

  "github.com/prometheus/client_golang/prometheus"
  log "github.com/sirupsen/logrus"
)

const (
  MemoryName      = "Memory"
  MemoryUrlSuffix = "/restconf/data/sonusSystem:system/memoryUtilCurrentStatistics"
)

var MemoryMetric = lib.SonusMetric{
  Name:       MemoryName,
  Processor:  processMemory,
  URLGetter:  getMemoryUrl,
  APIMetrics: MemoryMetrics,
  Repetition: lib.RepeatNone,
}

func getMemoryUrl(ctx lib.MetricContext) string {
  return ctx.APIBase + MemoryUrlSuffix
}
/*
ceName	The host name is this server.
	portName	The name of physical port.
	state	The state of the physical port
	faultState	The fault state of the physical port
	linkState	The link state of the physical port
	failures	The current number failures of this port monitor
	linkFailures	The current number of link failures of this port monitor

  CE Name	The server name.
  Prg Name	The name of the port redundancy group.
  Pm Name	The name of this port Monitor.
  Port Name	The name of physical port.
  Mac Address	The MAC address associated with the physical port.
  Role	
  The role of the physical port.
  
  Active
  No Role
  Standby
  State	
  The state of the physical port.
  
  Down
  Unknown
  up
  Fault State	
  The fault state of the physical port.
  
  Failed
  Restored
  Unknown
  Link State	
  The state of the link monitor configured on this port in the link detection group. The Link State is set to “Unknown” if link monitors are not configured.
  
  Failed
  Restored
  Unknown
  Failures	Current number failures of this port monitor.
  Link Failures	The current number of link failures of this port monitor.

var MemoryMetrics = map[string]*prometheus.Desc{
  "Memory_Average": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "average"),
    "The average memory % utilization for this interval.",
    []string{"server"}, nil,
  ),
  "Memory_High": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "high"),
    "The high memory % utilization for this interval.",
    []string{"server"}, nil,
  ),
  "Memory_Low": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "low"),
    "The low memory % utilization for this interval.",
    []string{"server"}, nil,
  ),
  "Memory_AverageSwap": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "AverageSwap"),
    "The average swap % memory utilization for this interval.",
    []string{"server"}, nil,
  ),
  "Memory_HighSwap": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "HighSwap"),
    "The high swap % memory utilization for this interval.",
    []string{"server"}, nil,
  ),
  "Memory_LowSwap": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "memory", "LowSwap"),
    "The low swap memory % utilization for this interval.",
    []string{"server"}, nil,
  ),
}

func processMemory(ctx lib.MetricContext, xmlBody *[]byte) {
  var (
    errors        []*error
    memoryUnits = new(MemoryCollection)
  )

  err := xml.Unmarshal(*xmlBody, &memoryUnits)

  if err != nil {
    log.Errorf("Failed to deserialize memory XML: %v", err)
    errors = append(errors, &err)
    ctx.ResultChannel <- lib.MetricResult{Name: MemoryName, Success: false, Errors: errors}
    return
  }

  for _, memory := range memoryUnits.MemoryUtilCurrentStatistics /*powerSupplies.PowerSupplyStatus {
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_Average"], prometheus.GaugeValue, memory.Average, memory.CeName)
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_High"], prometheus.GaugeValue, memory.High, memory.CeName)
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_Low"], prometheus.GaugeValue, memory.Low, memory.CeName)
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_AverageSwap"], prometheus.GaugeValue, memory.AverageSwap, memory.CeName)
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_HighSwap"], prometheus.GaugeValue, memory.LowSwap, memory.CeName)
    ctx.MetricChannel <- prometheus.MustNewConstMetric(MemoryMetrics["Memory_LowSwap"], prometheus.GaugeValue, memory.HighSwap, memory.CeName)

  }

  log.Info("Memory Metrics collected")
  ctx.ResultChannel <- lib.MetricResult{Name: CpuName, Success: true}
}
  
type PortMonitorStatusCollection struct {
  PortMonitorStatuses []*PortMonitorStatus `xml:"memoryUtilCurrentStatistics,omitempty"`
}

type PortMonitorStatus struct {
  CeName  string    `xml:"ceName"`//http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 ceName
  PortName string   `xml:"portName"`
  State    string   `xml:"state"`
  FaultState     float64   `xml:"faultState"`
  LinkState     string   `xml:"linkState"`
  Failures     float64   `xml:"failures"`
  LinkFailures     float64   `xml:"linkFailures"`
}
*/
/*
collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <portMonitorStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:ORCA_SYSTEM="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01ta</ceName>
    <prgName>prg_mgt0</prgName>
    <pmName>mgt0</pmName>
    <portName>mgt0</portName>
    <macAddress>0:10:6b:5:59:de</macAddress>
    <role>noRole</role>
    <state>up</state>
    <faultState>restored</faultState>
    <linkState>unknown</linkState>
    <failures>0</failures>
    <linkFailures>0</linkFailures>
  </portMonitorStatus>
  <portMonitorStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:ORCA_SYSTEM="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01ta</ceName>
    <prgName>prg_mgt1</prgName>
    <pmName>mgt1</pmName>
    <portName>mgt1</portName>
    <macAddress>0:10:6b:5:59:df</macAddress>
    <role>noRole</role>
    <state>up</state>
    <faultState>restored</faultState>
    <linkState>unknown</linkState>
    <failures>0</failures>
    <linkFailures>0</linkFailures>
  </portMonitorStatus>
  <portMonitorStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:ORCA_SYSTEM="http://sonusnet.com/ns/mibs/SONUS-ORCA-SYSTEM/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01tb</ceName>
    <prgName>prg_pkt3</prgName>
    <pmName>pkt3</pmName>
    <portName>pkt3</portName>
    <macAddress>0:10:6b:5:5a:ab</macAddress>
    <role>noRole</role>
    <state>down</state>
    <faultState>unknown</faultState>
    <linkState>unknown</linkState>
    <failures>1</failures>
    <linkFailures>0</linkFailures>
  </portMonitorStatus>

	ceName	The host name is this server.
	portName	The name of physical port.
	state	The state of the physical port
	faultState	The fault state of the physical port
	linkState	The link state of the physical port
	failures	The current number failures of this port monitor
	linkFailures	The current number of link failures of this port monitor
*/
