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

func processMemory(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
  var (
    errors        []*error
    memoryUnits = new(MemoryCollection)
  )

  err := xml.Unmarshal(*xmlBody, &memoryUnits)

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <memoryUtilCurrentStatistics xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01ta</ceName>
    <average>39</average>
    <high>39</high>
    <low>39</low>
    <averageSwap>0</averageSwap>
    <highSwap>0</highSwap>
    <lowSwap>0</lowSwap>
  </memoryUtilCurrentStatistics>
........

*/

  if err != nil {
    log.Errorf("Failed to deserialize memory XML: %v", err)
    errors = append(errors, &err)
    ctx.ResultChannel <- lib.MetricResult{Name: MemoryName, Success: false, Errors: errors}
    return
  }

  for _, memory := range memoryUnits.MemoryUtilCurrentStatistics /*powerSupplies.PowerSupplyStatus*/ {
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_Average"], prometheus.GaugeValue, memory.Average, memory.CeName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_High"], prometheus.GaugeValue, memory.High, memory.CeName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_Low"], prometheus.GaugeValue, memory.Low, memory.CeName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_AverageSwap"], prometheus.GaugeValue, memory.AverageSwap, memory.CeName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_HighSwap"], prometheus.GaugeValue, memory.LowSwap, memory.CeName))
    ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(MemoryMetrics["Memory_LowSwap"], prometheus.GaugeValue, memory.HighSwap, memory.CeName))

  }

  log.Info("Memory Metrics collected")
  ctx.ResultChannel <- lib.MetricResult{Name: MemoryName, Success: true}
}
  
type MemoryCollection struct {
  MemoryUtilCurrentStatistics []*MemoryStatus `xml:"memoryUtilCurrentStatistics,omitempty"`
}

type MemoryStatus struct {
  CeName  string    `xml:"ceName"`
  Average float64   `xml:"average"`
  High    float64   `xml:"high"`
  Low     float64   `xml:"low"`
  AverageSwap     float64   `xml:"averageSwap"`
  HighSwap     float64   `xml:"highSwap"`
  LowSwap     float64   `xml:"lowSwap"`

}



