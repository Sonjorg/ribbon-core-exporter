package metrics

// HA server sync status details

import (
  "encoding/xml"

  "core-exporter/lib"

  "github.com/prometheus/client_golang/prometheus"
  log "github.com/sirupsen/logrus"
)

const (
  SyncStatusName      = "SyncStatus"
  SyncStatusUrlSuffix = "/restconf/data/sonusSystem:system/syncStatus"
)

var SyncStatusMetric = lib.SonusMetric{
  Name:       SyncStatusName,
  Processor:  processSyncStatus,
  URLGetter:  getSyncStatusUrl,
  APIMetrics: SyncStatusMetrics,
  Repetition: lib.RepeatNone,
}

func getSyncStatusUrl(ctx lib.MetricContext) string {
  return ctx.APIBase + SyncStatusUrlSuffix
}


var SyncStatusMetrics = map[string]*prometheus.Desc{
  "Status": prometheus.NewDesc(
    prometheus.BuildFQName("ribbon", "system", "sync_status_details"),
    "Indicates the server data synchronization state 1:Sync completed, 0:not completed",
    []string{"syncModule","system"}, nil,
  ),  
}

func processSyncStatus(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
  var (
    errors        []*error
    syncStatus = new(SyncStatusCollection)
  )

  err := xml.Unmarshal(*xmlBody, &syncStatus)

  if err != nil {
    log.Errorf("Failed to deserialize SyncStatus XML: %v", err)
    errors = append(errors, &err)
    ctx.ResultChannel <- lib.MetricResult{Name: SyncStatusName, Success: false, Errors: errors}
    return
  }
	for _, name := range system {

   for _, status := range syncStatus.SyncStatuses {
     ctx.MetricChannel <- prometheus.MustNewConstMetric(SyncStatusMetrics["Status"], prometheus.GaugeValue, syncStatusfunc(status.Status), status.SyncModule, name)
   }
  }
  log.Info("Sync status metrics collected")
  ctx.ResultChannel <- lib.MetricResult{Name: SyncStatusName, Success: true}
}
  

/*
<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <syncStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <syncModule>Policy Data</syncModule>
    <status>syncCompleted</status>
  </syncStatus>
*/


type SyncStatusCollection struct {
  SyncStatuses []*SyncStatus `xml:"syncStatus,omitempty"`
}

type SyncStatus struct {
  SyncModule  string    `xml:"syncModule"`
  Status   string    `xml:"status"`
}

func syncStatusfunc(status string) float64{
  if (status == "syncCompleted"){
    return 1
   }else {
      return 0}
}
