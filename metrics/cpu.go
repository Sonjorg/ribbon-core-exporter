
//CPU usage metrics...

  package metrics

  import (
    "encoding/xml"
  
    "core-exporter/lib"
  
    "github.com/prometheus/client_golang/prometheus"
    log "github.com/sirupsen/logrus"
  )
  
  const (
    CpuName      = "CPU"
    CpuUrlSuffix = "/restconf/data/sonusSystem:system/cpuUtilCurrentStatistics"
  )
  
  var CpuMetric = lib.SonusMetric{
    Name:       CpuName,
    Processor:  processCpu,
    URLGetter:  getCpuUrl,
    APIMetrics: CpuMetrics,
    Repetition: lib.RepeatNone,
  }
  
  func getCpuUrl(ctx lib.MetricContext) string {
    return ctx.APIBase + CpuUrlSuffix
  }

  var CpuMetrics = map[string]*prometheus.Desc{
    "Cpu_Average": prometheus.NewDesc(
      prometheus.BuildFQName("ribbon", "cpu", "average"),
      "The average cpu % utilization for this interval.",
      []string{"server", "cpuID"}, nil,
    ),
    "Cpu_High": prometheus.NewDesc(
      prometheus.BuildFQName("ribbon", "cpu", "high"),
      "The high cpu % utilization for this interval.",
      []string{"server", "cpuID"}, nil,
    ),
    "Cpu_Low": prometheus.NewDesc(
      prometheus.BuildFQName("ribbon", "cpu", "low"),
      "The low cpu % utilization for this interval.",
      []string{"server", "cpuID"}, nil,
    ),
  }
  
  func processCpu(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
    var (
      errors        []*error
      cpus = new(CpuCollection)
    )
  
    err := xml.Unmarshal(*xmlBody, &cpus)
  
    if err != nil {
      log.Errorf("Failed to deserialize cpu XML: %v", err)
      errors = append(errors, &err)
      ctx.ResultChannel <- lib.MetricResult{Name: CpuName, Success: false, Errors: errors}
      return
    }

    /*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <cpuUtilCurrentStatistics xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <ceName>NOGJHDO-SBC-01ta</ceName>
    <cpu>1</cpu>
    <average>0</average>
    <high>2</high>
    <low>0</low>
  </cpuUtilCurrentStatistics>
 */
  
    for _, cpu := range cpus.CpuUtilCurrentStatistics {
      ctx.MetricChannel <- prometheus.MustNewConstMetric(CpuMetrics["Cpu_Average"], prometheus.GaugeValue, cpu.Average, cpu.CeName, cpu.CpuID)
      ctx.MetricChannel <- prometheus.MustNewConstMetric(CpuMetrics["Cpu_High"], prometheus.GaugeValue, cpu.High, cpu.CeName, cpu.CpuID)
      ctx.MetricChannel <- prometheus.MustNewConstMetric(CpuMetrics["Cpu_Low"], prometheus.GaugeValue, cpu.Low, cpu.CeName, cpu.CpuID)
    }
  
    log.Info("CPU Metrics collected")
    ctx.ResultChannel <- lib.MetricResult{Name: CpuName, Success: true}
  }
    
  type CpuCollection struct {
    CpuUtilCurrentStatistics []*CpuStatus `xml:"cpuUtilCurrentStatistics,omitempty"`
  }
  
  type CpuStatus struct {
    CeName  string    `xml:"ceName"`
    CpuID   string    `xml:"cpu"`
    Average float64   `xml:"average"`
    High    float64   `xml:"high"`
    Low     float64   `xml:"low"`
  }

