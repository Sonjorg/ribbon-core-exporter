//Call count data...

package metrics

import (

	"core-exporter/lib"
	"encoding/xml"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	CallCountName      = "GlobalCallCountStatus"
	CallCountUrlSuffix = "/restconf/data/sonusGlobal:global/sonusActiveCall:callCountStatus"
)

var CallCountMetric = lib.SonusMetric{
	Name:       CallCountName,
	Processor:  processCallCount,
	URLGetter:  getCallCountUrl,
	APIMetrics: CallCountMetrics,
	Repetition: lib.RepeatNone,
}

func getCallCountUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + CallCountUrlSuffix
}


var CallCountMetrics = map[string]*prometheus.Desc{
	"callAttempts": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "callAttempts"),
		"callAttempts",
		[]string{"key", "systemName"}, nil,
	),
	"callCompletions": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "callCompletions"),
		"callCompletions",
		[]string{"key", "systemName"}, nil,
	),
	"activeCalls": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "activeCalls"),
		"activeCalls",
		[]string{"key", "systemName"}, nil,
	),
	"stableCalls": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "stableCalls"),
		"stableCalls.",
		[]string{"key", "systemName"}, nil,
	),
	"callUpdates": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "callUpdates"),
		"callUpdates.",
		[]string{"key", "systemName"}, nil,
	),
	"activeCallsNonUser": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "activeCallsNonUser"),
		"activeCallsNonUser.",
		[]string{"key", "systemName"}, nil,
	),
	"stableCallsNonUser": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "stableCallsNonUser"),
		"stableCallsNonUser.",
		[]string{"key", "systemName"}, nil,
	), "totalCalls": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "totalCalls"),
		"totalCalls.",
		[]string{"key", "systemName"}, nil,
	), "totalCallsNonUser": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "totalCallsNonUser"),
		"totalCallsNonUser.",
		[]string{"key", "systemName"}, nil,
	), "totalCallsEmergEstablishing": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "totalCallsEmergEstablishing"),
		"totalCallsEmergEstablishing.",
		[]string{"key", "systemName"}, nil,
	), "totalCallsEmergStable": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "global", "totalCallsEmergStable"),
		"totalCallsEmergStable.",
		[]string{"key", "systemName"}, nil,
	),
}

func processCallCount(ctx lib.MetricContext, xmlBody *[]byte,system []string) {
	var (
		errors     []*error
		callCounts = new(CallCountStatusCollection)
	)

  if len(*xmlBody) == 0 {
		ctx.ResultChannel <- lib.MetricResult{Name: CallCountName, Success: true}
		return
	}
	err := xml.Unmarshal(*xmlBody, &callCounts)
if len(callCounts.CallCountStatuses)==0{
  return
}

	if err != nil {
		log.Errorf("Failed to deserialize memory XML: %v", err)
		errors = append(errors, &err)
		ctx.ResultChannel <- lib.MetricResult{Name: CallCountName, Success: false, Errors: errors}
		return
	}

	for _, name := range system {
		for _, callCount := range callCounts.CallCountStatuses {
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["callAttempts"], prometheus.GaugeValue, callCount.CallAttempts, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["callCompletions"], prometheus.GaugeValue, callCount.CallCompletions, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["activeCalls"], prometheus.GaugeValue, callCount.ActiveCalls, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["stableCalls"], prometheus.GaugeValue, callCount.StableCalls, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["callUpdates"], prometheus.GaugeValue, callCount.CallUpdates, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["activeCallsNonUser"], prometheus.GaugeValue, callCount.ActiveCallsNonUser, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["stableCallsNonUser"], prometheus.GaugeValue, callCount.StableCallsNonUser, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["totalCalls"], prometheus.GaugeValue, callCount.TotalCalls, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["totalCallsNonUser"], prometheus.GaugeValue, callCount.TotalCallsNonUser, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["totalCallsEmergEstablishing"], prometheus.GaugeValue, callCount.TotalCallsEmergEstablishing, callCount.Key, name))
			ctx.MetricArray = append(ctx.MetricArray,prometheus.MustNewConstMetric(CallCountMetrics["totalCallsEmergStable"], prometheus.GaugeValue, callCount.TotalCallsEmergStable, callCount.Key, name))
		}
	}

	log.Info("Call Count Metrics collected")
	ctx.ResultChannel <- lib.MetricResult{Name: CallCountName, Success: true}
}


/*

<collection xmlns="http://tail-f.com/ns/restconf/collection/1.0">
  <callCountStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-ACTIVE-CALL/1.0"  xmlns:ACTIVE_CALL="http://sonusnet.com/ns/mibs/SONUS-ACTIVE-CALL/1.0"  xmlns:GLOBAL_OBJECTS="http://sonusnet.com/ns/mibs/SONUS-GLOBAL-OBJECTS/1.0">
    <key>all</key>
    <callAttempts>1</callAttempts>
    <callCompletions>0</callCompletions>
    <activeCalls>0</activeCalls>
    <stableCalls>0</stableCalls>
    <callUpdates>0</callUpdates>
    <activeCallsNonUser>0</activeCallsNonUser>
    <stableCallsNonUser>0</stableCallsNonUser>
    <totalCalls>0</totalCalls>
    <totalCallsNonUser>0</totalCallsNonUser>
    <totalCallsEmergEstablishing>0</totalCallsEmergEstablishing>
    <totalCallsEmergStable>0</totalCallsEmergStable>
  </callCountStatus>
</collection>
*/

type CallCountStatusCollection struct {
	CallCountStatuses []*CallCountStatus `xml:"callCountStatus,omitempty"`
}

type CallCountStatus struct {
	Key                         string  `xml:"key"` 
	CallAttempts                float64 `xml:"callAttempts"`
	CallCompletions             float64 `xml:"callCompletions"`
	ActiveCalls                 float64 `xml:"activeCalls"`
	StableCalls                 float64 `xml:"stableCalls"`
	CallUpdates                 float64 `xml:"callUpdates"`
	ActiveCallsNonUser          float64 `xml:"activeCallsNonUser"`
	StableCallsNonUser          float64 `xml:"stableCallsNonUser"`
	TotalCalls                  float64 `xml:"totalCalls"`
	TotalCallsNonUser           float64 `xml:"totalCallsNonUser"`
	TotalCallsEmergEstablishing float64 `xml:"totalCallsEmergEstablishing"`
	TotalCallsEmergStable       float64 `xml:"totalCallsEmergStable"`
}



