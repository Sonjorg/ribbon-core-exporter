package metrics

// software upgrade status


  import (
    "encoding/xml"
  
    "core-exporter/lib"
  
    "github.com/prometheus/client_golang/prometheus"
    log "github.com/sirupsen/logrus"
  )
  
  const (
    SoftwareUpgradeName      = "ServerSoftWareUpgradeStatus"
    SoftwareUpgradeUrlSuffix = "/restconf/data/sonusSystem:system/serverSoftwareUpgradeStatus"
  )
  
  var SoftwareUpgradeMetric = lib.SonusMetric{
    Name:       SoftwareUpgradeName,
    Processor:  processSoftwareUpgrade,
    URLGetter:  getSoftwareUpgradeUrl,
    APIMetrics: SoftwareUpgradeMetrics,
    Repetition: lib.RepeatNone,
  }
  
  func getSoftwareUpgradeUrl(ctx lib.MetricContext) string {
    return ctx.APIBase + SoftwareUpgradeUrlSuffix
  }

  var SoftwareUpgradeMetrics = map[string]*prometheus.Desc{
    "upgradeStatus": prometheus.NewDesc(
      prometheus.BuildFQName("ribbon", "system", "SoftwareUpgrade_status"),
      "Indicates the current status of the server: 0=upgraded, 1=upgrading, 2=pendingUpgrade",
      []string{"server"}, nil,
    ),
  }
  
  func processSoftwareUpgrade(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
    var (
      errors        []*error
      statuses = new(ServerSoftwareUpgradeStatusCollection)
    )
  
    err := xml.Unmarshal(*xmlBody, &statuses)
  
    if err != nil {
      log.Errorf("Failed to deserialize cpu XML: %v", err)
      errors = append(errors, &err)
      ctx.ResultChannel <- lib.MetricResult{Name: SoftwareUpgradeName, Success: false, Errors: errors}
      return
    }
  
    for _, status := range statuses.ServerSoftwareUpgradeStatuses {
      ctx.MetricChannel <- prometheus.MustNewConstMetric(SoftwareUpgradeMetrics["upgradeStatus"], prometheus.GaugeValue, convertUpgradeStatusToNum(status.UpgradeStatus), status.Server)
    }
  
    log.Info("SoftwareUpgrade Metrics collected")
    ctx.ResultChannel <- lib.MetricResult{Name: SoftwareUpgradeName, Success: true}
  }
   
    /*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <serverSoftwareUpgradeStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <name>NOGJHDO-SBC-01ta</name>
    <upgradeStatus>upgraded</upgradeStatus>
  </serverSoftwareUpgradeStatus>
  <serverSoftwareUpgradeStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <name>NOGJHDO-SBC-01tb</name>
    <upgradeStatus>upgraded</upgradeStatus>
  </serverSoftwareUpgradeStatus>
</collection>
*/

  type ServerSoftwareUpgradeStatusCollection struct {
    ServerSoftwareUpgradeStatuses []*ServerSoftwareUpgradeStatus `xml:"serverSoftwareUpgradeStatus,omitempty"`
  }
  
  type ServerSoftwareUpgradeStatus struct {
    Server         string    `xml:"name"`
    UpgradeStatus  string    `xml:"upgradeStatus"`
  }
  
  func convertUpgradeStatusToNum(status string)float64{
    switch status {
    case "upgraded":
      return 0
    case "upgrading":
      return 1
    case "pendingUpgrade":
      return 2
    }
    return 0
  }

