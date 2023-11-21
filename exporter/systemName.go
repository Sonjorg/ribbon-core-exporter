package exporter

import (
	"encoding/xml"
	//"strconv"
	//"strings"
	//"core-exporter/exporter"
	"net/http"
	"crypto/tls"

	"core-exporter/lib"
	"core-exporter/config"

	//"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	systemName      = "systemVariable"
	systemUrlSuffix = "/restconf/data/sonusSystem:system/admin"
)


func getSystemUrl(ctx lib.MetricContext) string {
	return ctx.APIBase + systemUrlSuffix
}


func ProcessSystemName(url string) []string{
	var (
		errors []*error
		system   = new(AdminCollection)
		systemName []string
	)
	cfg := config.GetConfig()
	//getSystemUrl(ctx)
	httpTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	httpClient := &http.Client{Transport: httpTransport, Timeout: cfg.APITimeout}

	xmlbody,err := doHTTPRequest(httpClient, url, cfg.APIUser, cfg.APIPass)
	err = xml.Unmarshal(*xmlbody.body, &system)
	//log.Info(cfg.APIUser,*xmlbody.body)

	if err != nil {
		log.Errorf("Failed to deserialize system name XML: %v", err)
		errors = append(errors, &err)

		systemName = append(systemName, "Error fetching sysName")
		return systemName
	}
	for _, sys := range system.Admins {
		if err != nil {
			log.Errorf("Failed to fetch system name: %s",err)
			errors = append(errors, &err)
			break
		}
		 systemName = append(systemName, sys.ActualName)
		// println(systemName)
	}

	log.Info("system name collected")

	return systemName
}



type AdminCollection struct {
	Admins []*Admin `xml:"admin,omitempty"`
}

type Admin struct {
	Name string `xml:"name"`
	ActualName string `xml:"actualSystemName"`
}
/*
<admin xmlns="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0"  xmlns:SYS="http://sonusnet.com/ns/mibs/SONUS-SYSTEM-MIB/1.0">
  <name>NOGJHDO-SBC-01t</name>
  <actualSystemName>NOGJHDO-SBC-01t</actualSystemName>
  <haMode>haMode1to1</haMode>
  <auditLogState>enabled</auditLogState>
  <dspMismatchAction>preserveRedundancy</dspMismatchAction>
  <passwordRules>
    <minimumLength>8</minimumLength>
    <minimumNumberOfUppercaseChars>1</minimumNumberOfUppercaseChars>
    <minimumNumberOfLowercaseChars>1</minimumNumberOfLowercaseChars>
    <minimumNumberOfDigits>1</minimumNumberOfDigits>
    <minimumNumberOfOtherChars>1</minimumNumberOfOtherChars>
    <passwordHistoryDepth>4</passwordHistoryDepth>
    <maximumRepeatingCharsCount>3</maximumRepeatingCharsCount>
    <minimumDiffWithOldPassword>4</minimumDiffWithOldPassword>
  </passwordRules>
  <accountManagement>
    <sessionIdleTimeout>
      <idleTimeout>30</idleTimeout>
    </sessionIdleTimeout>
    <maxSessions>5</maxSessions>
  </accountManagement>
  <fips-140-2>
    <mode>disabled</mode>
  </fips-140-2>
  <ema>
    <enableREST>enabled</enableREST>
    <enableCoreEMA>enabled</enableCoreEMA>
    <enableTS>enabled</enableTS>
  </ema>
  <sshPublicKeyAuthenticationEnabled>true</sshPublicKeyAuthenticationEnabled>
</admin>
*/
