package exporter

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"core-exporter/lib"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func getServerStatusURL(apiBase string) string {
	return fmt.Sprintf("%s/restconf/data/sonusSystem:system/serverStatus", apiBase)
}

func (a *addressContext) getZoneStatusURL(ctx lib.MetricContext) string {
	return fmt.Sprintf("%s/restconf/data/sonusAddressContext:addressContext=%s/sonusZone:zoneStatus", ctx.APIBase, a.Name)

}

func (a *addressContext) getIPInterfaceGroupURL(ctx lib.MetricContext) string {
	return fmt.Sprintf("%s/restconf/data/sonusAddressContext:addressContext=%s/sonusIpInterface:ipInterfaceGroup", ctx.APIBase, a.Name)
}
var serverStatusMetrics = map[string]*prometheus.Desc{
	"System_Redundancy_Role": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "system", "redundancy_role"),
		"Current role of server. 1 = active",
		[]string{"server", "role_name"}, nil,
	),
	"System_Sync_Status": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "system", "sync_status"),
		"Current synchronization status. 1 = syncCompleted",
		[]string{"server", "status_name"}, nil,
	),
	"System_Uptime": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "system", "uptime"),
		"Current uptime of server, in seconds",
		[]string{"server", "type"}, nil,
	),
	"System_Info": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "system", "info"),
		"Current status of server",
		[]string{"server", "HwType", "restart", "platformVersion", "applicationVersion", "serial"}, nil,
	),
}

var zoneStatusMetrics = map[string]*prometheus.Desc{
	"Zone_Total_Calls_Configured": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "zone", "channels_configured"),
		"Number of configured calls per zone",
		[]string{"addresscontext", "zone"}, nil,
	),
	"Zone_Usage_Total": prometheus.NewDesc(
		prometheus.BuildFQName("ribbon", "zone", "activeCalls"),
		"Total active calls per per zone",
		[]string{"direction", "addresscontext", "zone"}, nil,
	),
}

type (
	serverStatusCollection struct {
		ServerStatus []*serverStatus `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 serverStatus"`
	}


/*

<serverStatus xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
    <name>NOGJHDO-SBC-01ta</name>
    <hwType>SBC 5400</hwType>
    <serialNum>2214220019</serialNum>
    <partNum>821-00504</partNum>
    <platformVersion>V11.01.00R001</platformVersion>
    <applicationVersion>V11.01.00R001</applicationVersion>
    <mgmtRedundancyRole>active</mgmtRedundancyRole>
    <upTime>99 Days 03:49:54</upTime>
    <applicationUpTime>99 Days 03:44:55</applicationUpTime>
    <lastRestartReason>systemRestart</lastRestartReason>
    <syncStatus>syncCompleted</syncStatus>
    <daughterBoardPresent>false</daughterBoardPresent>
    <currentTime>2023/11/03 15:17:29 </currentTime>
    <pktPortSpeed>speed1Gbps</pktPortSpeed>
    <actualCeName>NOGJHDO-SBC-01ta</actualCeName>
    <hwSubType>5400</hwSubType>
    <fingerprint></fingerprint>
  </serverStatus>

*/



	serverStatus struct {
		Name                     string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 name"`
		SerialNum                string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 serialNum"`
		ManagementRedundancyRole string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 mgmtRedundancyRole"`
		Uptime                   string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 upTime"`
		ApplicationUptime        string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 applicationUpTime"`
		SyncStatus               string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 syncStatus"`
		LastRestartReason        string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 lastRestartReason"`
		PlatformVersion          string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 platformVersion"`
		ApplicationVersion       string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 applicationVersion"`
		HwType          		 string `xml:"http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0 hwType"`
					
	}

	serverUptimeType uint8
)

const (
	serverOSUptime serverUptimeType = iota
	serverAppUptime
)

var uptimeRegex = regexp.MustCompile(`^(\d{1,5}) Days (\d{2}):(\d{2}):(\d{2})$`)

func (s serverStatus) parseUptime(upType serverUptimeType) float64 {
	var (
		uptime    time.Duration
		uptimeStr string
	)

	switch upType {
	case serverOSUptime:
		uptimeStr = s.Uptime
	case serverAppUptime:
		uptimeStr = s.ApplicationUptime
	}

	uptimeFields := uptimeRegex.FindStringSubmatch(uptimeStr)

	if len(uptimeFields) == 5 {
		daysInt, _ := strconv.ParseInt(uptimeFields[1], 10, 32)
		hoursInt, _ := strconv.ParseInt(uptimeFields[2], 10, 32)
		minutesInt, _ := strconv.ParseInt(uptimeFields[3], 10, 32)
		secondsInt, _ := strconv.ParseInt(uptimeFields[4], 10, 32)

		uptime += time.Duration(daysInt) * time.Hour * 24
		uptime += time.Duration(hoursInt) * time.Hour
		uptime += time.Duration(minutesInt) * time.Minute
		uptime += time.Duration(secondsInt) * time.Second

		return uptime.Seconds()
	} else {
		log.Errorf("Unable to match uptime %q with regex.", uptimeStr)
		return 0
	}
}

func (s serverStatus) mgmtRedunRoleToFloat() float64 {
	switch s.ManagementRedundancyRole {
	case "active":
		return 1
	default:
		return 0
	}
}

func (s serverStatus) syncStatusToFloat() float64 {
	switch s.SyncStatus {
	case "syncCompleted":
		return 1
	default:
		return 0
	}
}

type addressContext struct {
	Name              string
	Zones             []*zoneStatus       `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 zoneStatus"`
	IPInterfaceGroups []*ipInterfaceGroup `xml:"http://sonusnet.com/ns/mibs/SONUS-GEN2-IP-INTERFACE/1.0 ipInterfaceGroup"`
}

type zoneStatus struct {
	Name                 string  `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 name"`
	TotalCallsAvailable  float64 `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 totalCallsAvailable"`
	InboundCallsUsage    float64 `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 inboundCallsUsage"`
	OutboundCallsUsage   float64 `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 outboundCallsUsage"`
	TotalCallsConfigured float64 `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 totalCallsConfigured"`
	ActiveSipRegCount    float64 `xml:"http://sonusnet.com/ns/mibs/SONUS-ZONE/1.0 activeSipRegCount"`
}

type ipInterfaceGroup struct {
	Name         string `xml:"name"`
	IpInterfaces struct {
		Name string `xml:"name"`
	} `xml:"http://sonusnet.com/ns/mibs/SONUS-GEN2-IP-INTERFACE/1.0 ipInterface"`
}

func processServerStatus(xmlBody *[]byte, ch chan<- prometheus.Metric) error {
	var serverStatuses = new(serverStatusCollection)

	err := xml.Unmarshal(*xmlBody, &serverStatuses)
	if err != nil {
		log.Errorf("Failed to deserialize serverStatus XML: %v", err)
		return err
	}

	for _, server := range serverStatuses.ServerStatus {
		ch <- prometheus.MustNewConstMetric(serverStatusMetrics["System_Redundancy_Role"], prometheus.GaugeValue, server.mgmtRedunRoleToFloat(), server.Name, server.ManagementRedundancyRole)
		ch <- prometheus.MustNewConstMetric(serverStatusMetrics["System_Sync_Status"], prometheus.GaugeValue, server.syncStatusToFloat(), server.Name, server.SyncStatus)
		ch <- prometheus.MustNewConstMetric(serverStatusMetrics["System_Uptime"], prometheus.CounterValue, server.parseUptime(serverOSUptime), server.Name, "os")
		ch <- prometheus.MustNewConstMetric(serverStatusMetrics["System_Uptime"], prometheus.CounterValue, server.parseUptime(serverAppUptime), server.Name, "application")
		ch <- prometheus.MustNewConstMetric(serverStatusMetrics["System_Info"], prometheus.GaugeValue, 1, server.Name, server.HwType, server.LastRestartReason, server.PlatformVersion, server.ApplicationVersion, server.SerialNum)
	}  // 
	log.Info("Server Status and Metrics collected")
	return nil
}

func processZones(addressContext *addressContext, xmlBody *[]byte, ch chan<- prometheus.Metric) error {
	err := xml.Unmarshal(*xmlBody, &addressContext)
	if err != nil {
		log.Errorf("Failed to deserialize zoneStatus XML: %v", err)
		return err
	}

	for _, zone := range addressContext.Zones {
		ch <- prometheus.MustNewConstMetric(zoneStatusMetrics["Zone_Total_Calls_Configured"], prometheus.GaugeValue, zone.TotalCallsConfigured, addressContext.Name, zone.Name)
		ch <- prometheus.MustNewConstMetric(zoneStatusMetrics["Zone_Usage_Total"], prometheus.GaugeValue, zone.InboundCallsUsage, "inbound", addressContext.Name, zone.Name)
		ch <- prometheus.MustNewConstMetric(zoneStatusMetrics["Zone_Usage_Total"], prometheus.GaugeValue, zone.OutboundCallsUsage, "outbound", addressContext.Name, zone.Name)
	}
	log.Info("Zone Status and Metrics collected")
	return nil
}

func processIPInterfaceGroups(addressContext *addressContext, xmlBody *[]byte) error {
	err := xml.Unmarshal(*xmlBody, &addressContext)
	if err != nil {
		log.Errorf("Failed to deserialize ipInterfaceGroup XML: %v", err)
		return err
	}
	//fmt.Println("Interface groups: ",addressContext.IPInterfaceGroups)
 //   log.Info("IpInterfaceGroup: %v collected")
	return nil
}
