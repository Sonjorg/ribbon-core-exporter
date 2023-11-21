# Metrics

Below are an example of the metrics as exposed by this exporter.
Its a limited view and many collections consist of more metrics

## DSP Statistics

```
# HELP ribbon_dsp_codec_utilization Codec utilization, in percent
# TYPE ribbon_dsp_codec_utilization gauge
.....
ribbon_dsp_codec_utilization{codec="G.711",system="NOGJHDO-SBC-01t"} 0
ribbon_dsp_codec_utilization{codec="G.711 Silence Suppression",system="NOGJHDO-SBC-01t"} 0
ribbon_dsp_codec_utilization{codec="G.711 Silence Suppression V8",system="NOGJHDO-SBC-01t"} 0
....


# HELP ribbon_dsp_compression_utilization Compression resource utilization, in percent
# TYPE ribbon_dsp_compression_utilization gauge
ribbon_dsp_compression_utilization{system="NOGJHDO-SBC-01t"} 0

# HELP ribbon_dsp_resources_total Total compression resources
# TYPE ribbon_dsp_resources_total gauge
ribbon_dsp_resources_total{system="NOGJHDO-SBC-01t"} 0

# HELP ribbon_dsp_resources_used Usage of DSP resources per slot
# TYPE ribbon_dsp_resources_used gauge
ribbon_dsp_resources_used{slot="1",system="NOGJHDO-SBC-01t"} 0
ribbon_dsp_resources_used{slot="2",system="NOGJHDO-SBC-01t"} 0
ribbon_dsp_resources_used{slot="3",system="NOGJHDO-SBC-01t"} 0
ribbon_dsp_resources_used{slot="4",system="NOGJHDO-SBC-01t"} 0
```

## Fan Status

```
# HELP ribbon_fan_speed Current speed of fans, in RPM
# TYPE ribbon_fan_speed gauge
ribbon_fan_speed{fanID="FAN_CTR/TACH_F",server="NOGJHDO-SBC-01ta"} 4352
ribbon_fan_speed{fanID="FAN_CTR/TACH_F",server="NOGJHDO-SBC-01tb"} 4352
....

```

## IP Interface Statistics

```
# HELP ribbon_ipinterface_media_streams Number of media streams currently on ipInterfaceGroup
# TYPE ribbon_ipinterface_media_streams gauge
ribbon_ipinterface_media_streams{name="IPIF1"} 0
ribbon_ipinterface_media_streams{name="IPIF4"} 0

# HELP ribbon_ipinterface_rxbandwidth Receive bandwidth in use on interface, in bytes per second
# TYPE ribbon_ipinterface_rxbandwidth gauge
ribbon_ipinterface_rxbandwidth{name="IPIF1"} 0
ribbon_ipinterface_rxbandwidth{name="IPIF4"} 0


# HELP ribbon_ipinterface_rxpackets Number of packets received on ipInterfaceGroup
# TYPE ribbon_ipinterface_rxpackets counter
ribbon_ipinterface_rxpackets{name="IPIF1"} 1.985825e+06
ribbon_ipinterface_rxpackets{name="IPIF4"} 1.393682e+06


# HELP ribbon_ipinterface_status Current status of ipInterfaceGroup
# TYPE ribbon_ipinterface_status gauge
ribbon_ipinterface_status{name="IPIF1",status_text="resAllocated"} 0
ribbon_ipinterface_status{name="IPIF4",status_text="resAllocated"} 0


# HELP ribbon_ipinterface_txbandwidth Transmit bandwidth in use on interface, in bytes per second
# TYPE ribbon_ipinterface_txbandwidth gauge
ribbon_ipinterface_txbandwidth{name="IPIF1"} 0
ribbon_ipinterface_txbandwidth{name="IPIF4"} 0


# HELP ribbon_ipinterface_txpackets Number of packets transmitted on ipInterfaceGroup
# TYPE ribbon_ipinterface_txpackets counter
ribbon_ipinterface_txpackets{name="IPIF1"} 1.985861e+06
ribbon_ipinterface_txpackets{name="IPIF4"} 1.393779e+06

```

## Power Supplies

```
# HELP ribbon_powersupply_powerfault Is there a power fault, per supply
# TYPE ribbon_powersupply_powerfault gauge
ribbon_powersupply_powerfault{powerSupplyID="PS_BOT",server="NOGJHDO-SBC-01ta"} 0
ribbon_powersupply_powerfault{powerSupplyID="PS_BOT",server="NOGJHDO-SBC-01tb"} 0
ribbon_powersupply_powerfault{powerSupplyID="PS_TOP",server="NOGJHDO-SBC-01ta"} 0
ribbon_powersupply_powerfault{powerSupplyID="PS_TOP",server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_powersupply_voltagefault Is there a voltage fault, per supply
# TYPE ribbon_powersupply_voltagefault gauge
ribbon_powersupply_voltagefault{powerSupplyID="PS_BOT",server="NOGJHDO-SBC-01ta"} 0
ribbon_powersupply_voltagefault{powerSupplyID="PS_BOT",server="NOGJHDO-SBC-01tb"} 0
ribbon_powersupply_voltagefault{powerSupplyID="PS_TOP",server="NOGJHDO-SBC-01ta"} 0
ribbon_powersupply_voltagefault{powerSupplyID="PS_TOP",server="NOGJHDO-SBC-01tb"} 0
```

## Memory Usage

```
# HELP ribbon_memory_AverageSwap The average swap % memory utilization for this interval.
# TYPE ribbon_memory_AverageSwap gauge
ribbon_memory_AverageSwap{server="NOGJHDO-SBC-01ta"} 0
ribbon_memory_AverageSwap{server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_memory_HighSwap The high swap % memory utilization for this interval.
# TYPE ribbon_memory_HighSwap gauge
ribbon_memory_HighSwap{server="NOGJHDO-SBC-01ta"} 0
ribbon_memory_HighSwap{server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_memory_LowSwap The low swap memory % utilization for this interval.
# TYPE ribbon_memory_LowSwap gauge
ribbon_memory_LowSwap{server="NOGJHDO-SBC-01ta"} 0
ribbon_memory_LowSwap{server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_memory_average The average memory % utilization for this interval.
# TYPE ribbon_memory_average gauge
ribbon_memory_average{server="NOGJHDO-SBC-01ta"} 39
ribbon_memory_average{server="NOGJHDO-SBC-01tb"} 17

# HELP ribbon_memory_high The high memory % utilization for this interval.
# TYPE ribbon_memory_high gauge
ribbon_memory_high{server="NOGJHDO-SBC-01ta"} 39
ribbon_memory_high{server="NOGJHDO-SBC-01tb"} 17

# HELP ribbon_memory_low The low memory % utilization for this interval.
# TYPE ribbon_memory_low gauge
ribbon_memory_low{server="NOGJHDO-SBC-01ta"} 39
ribbon_memory_low{server="NOGJHDO-SBC-01tb"} 17
```

## CPU Usage

```
# HELP ribbon_cpu_average The average cpu % utilization for this interval.
# TYPE ribbon_cpu_average gauge
ribbon_cpu_average{cpuID="1",server="NOGJHDO-SBC-01ta"} 1
ribbon_cpu_average{cpuID="1",server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_cpu_high The high cpu % utilization for this interval.
# TYPE ribbon_cpu_high gauge
ribbon_cpu_high{cpuID="1",server="NOGJHDO-SBC-01ta"} 3
ribbon_cpu_high{cpuID="1",server="NOGJHDO-SBC-01tb"} 2

# HELP ribbon_cpu_low The low cpu % utilization for this interval.
# TYPE ribbon_cpu_low gauge
ribbon_cpu_low{cpuID="1",server="NOGJHDO-SBC-01ta"} 1
ribbon_cpu_low{cpuID="1",server="NOGJHDO-SBC-01tb"} 0
```


## Disk Usage and Status

```
# HELP ribbon_disk_healthTest Pass or Fail indicating the overall health of the device. 1=passed
# TYPE ribbon_disk_healthTest gauge
ribbon_disk_healthTest{server="NOGJHDO-SBC-01ta"} 1
ribbon_disk_healthTest{server="NOGJHDO-SBC-01tb"} 1

# HELP ribbon_disk_size Capacity of the disk in the server
# TYPE ribbon_disk_size gauge
ribbon_disk_size{server="NOGJHDO-SBC-01ta"} 476
ribbon_disk_size{server="NOGJHDO-SBC-01tb"} 476

# HELP ribbon_disk_status Harddisk status that indicates if the disk is online/failed 1=online
# TYPE ribbon_disk_status gauge
ribbon_disk_status{productId="ATA WDC PC SA530 SDA",server="NOGJHDO-SBC-01ta"} 1
ribbon_disk_status{productId="ATA WDC PC SA530 SDA",server="NOGJHDO-SBC-01tb"} 1


# HELP ribbon_disk_free Indicates free hard disk space (KBytes)
# TYPE ribbon_disk_free gauge
ribbon_disk_free{Partition="/",server="NOGJHDO-SBC-01ta"} 4.386428e+06
ribbon_disk_free{Partition="/",server="NOGJHDO-SBC-01tb"} 4.42866e+06
....

# HELP ribbon_disk_used Indicates used hard disk (%)
# TYPE ribbon_disk_used gauge
ribbon_disk_used{Partition="/",server="NOGJHDO-SBC-01ta"} 57
ribbon_disk_used{Partition="/",server="NOGJHDO-SBC-01tb"} 57
...
```

## Server Status

```
# HELP ribbon_system_redundancy_role Current role of server. 1 = active
# TYPE ribbon_system_redundancy_role gauge
ribbon_system_redundancy_role{role_name="active",server="NOGJHDO-SBC-01ta"} 1
ribbon_system_redundancy_role{role_name="standby",server="NOGJHDO-SBC-01tb"} 0

# HELP ribbon_system_status Current status of server
# TYPE ribbon_system_status gauge
ribbon_system_status{platform="V11.01.00R001",restart="systemRestart",serial="2214220019",server="NOGJHDO-SBC-01ta"} 1
ribbon_system_status{platform="V11.01.00R001",restart="systemRestart",serial="2214220043",server="NOGJHDO-SBC-01tb"} 1

# HELP ribbon_system_sync_status Current synchronization status. 1 = syncCompleted
# TYPE ribbon_system_sync_status gauge
ribbon_system_sync_status{server="NOGJHDO-SBC-01ta",status_name="syncCompleted"} 1
ribbon_system_sync_status{server="NOGJHDO-SBC-01tb",status_name="syncCompleted"} 1

# HELP ribbon_system_uptime Current uptime of server, in seconds
# TYPE ribbon_system_uptime counter
ribbon_system_uptime{server="NOGJHDO-SBC-01ta",type="application"} 8.853288e+06
ribbon_system_uptime{server="NOGJHDO-SBC-01ta",type="os"} 8.853587e+06
ribbon_system_uptime{server="NOGJHDO-SBC-01tb",type="application"} 8.8524e+06
ribbon_system_uptime{server="NOGJHDO-SBC-01tb",type="os"} 8.852651e+06

```

## SIP ARS Status

```
# HELP ribbon_sipars_endpoint_status State of a sipArs monitored endpoint
# TYPE ribbon_sipars_endpoint_status gauge
ribbon_sipars_endpoint_status{endpoint_address="10.237.110.12",endpoint_port="5069",state_name="blacklisted",zone="PE_ICCS"} 1
```

## Trunk Groups

```
# HELP ribbon_TG_bytes Bandwidth in use by current calls
# TYPE ribbon_TG_bytes gauge

ribbon_TG_bytes{direction="inbound",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_bytes{direction="inbound",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_bytes{direction="outbound",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_bytes{direction="outbound",name="TG_ICCS337_T",zone="PE_ICCS"} 0

# HELP ribbon_TG_outbound_state State of outbound calls on the trunkgroup
# TYPE ribbon_TG_outbound_state gauge
ribbon_TG_outbound_state{name="TE_TELENOR_URD",zone="TE_URD"} 1
ribbon_TG_outbound_state{name="TG_ICCS337_T",zone="PE_ICCS"} 1

# HELP ribbon_TG_sip_req_recv Number of SIP requests received
# TYPE ribbon_TG_sip_req_recv counter
ribbon_TG_sip_req_recv{method="BYE",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_req_recv{method="BYE",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_sip_req_recv{method="CANCEL",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_req_recv{method="CANCEL",name="TG_ICCS337_T",zone="PE_ICCS"} 0
...

# HELP ribbon_TG_sip_req_sent Number of SIP requests sent
# TYPE ribbon_TG_sip_req_sent counter
ribbon_TG_sip_req_sent{method="BYE",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_req_sent{method="BYE",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_sip_req_sent{method="BYE (retrans)",name="TE_TELENOR_URD",zone="TE_URD"} 0
.....

# HELP ribbon_TG_sip_resp_recv Number of SIP responses received
# TYPE ribbon_TG_sip_resp_recv counter
ribbon_TG_sip_resp_recv{code="18x",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_resp_recv{code="18x",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_sip_resp_recv{code="1xx",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_resp_recv{code="1xx",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_sip_resp_recv{code="2xx",name="TE_TELENOR_URD",zone="TE_URD"} 0
....

# HELP ribbon_TG_sip_resp_sent Number of SIP responses sent
# TYPE ribbon_TG_sip_resp_sent counter
ribbon_TG_sip_resp_sent{code="18x",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_resp_sent{code="18x",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_sip_resp_sent{code="1xx",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_sip_resp_sent{code="1xx",name="TG_ICCS337_T",zone="PE_ICCS"} 0
.....

# HELP ribbon_TG_state State of the trunkgroup
# TYPE ribbon_TG_state gauge
ribbon_TG_state{name="TE_TELENOR_URD",zone="TE_URD"} 1
ribbon_TG_state{name="TG_ICCS337_T",zone="PE_ICCS"} 1

# HELP ribbon_TG_total_channels Number of configured channels
# TYPE ribbon_TG_total_channels gauge
ribbon_TG_total_channels{name="TE_TELENOR_URD",zone="TE_URD"} -1
ribbon_TG_total_channels{name="TG_ICCS337_T",zone="PE_ICCS"} -1

# HELP ribbon_TG_usage_total Number of active calls
# TYPE ribbon_TG_usage_total gauge
ribbon_TG_usage_total{direction="inbound",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_usage_total{direction="inbound",name="TG_ICCS337_T",zone="PE_ICCS"} 0
ribbon_TG_usage_total{direction="outbound",name="TE_TELENOR_URD",zone="TE_URD"} 0
ribbon_TG_usage_total{direction="outbound",name="TG_ICCS337_T",zone="PE_ICCS"} 0
```

## Exporter Metric Disposition

```
# HELP ribbon_exporter_metric_disposition Number of times each metric has succeeded or failed being collected
# TYPE ribbon_exporter_metric_disposition counter
ribbon_exporter_metric_disposition{name="CPU",successful="true"} 44
ribbon_exporter_metric_disposition{name="DSP",successful="true"} 22
ribbon_exporter_metric_disposition{name="DiskStatus",successful="true"} 22
ribbon_exporter_metric_disposition{name="Fan",successful="true"} 22
ribbon_exporter_metric_disposition{name="IPInterface",successful="true"} 88
ribbon_exporter_metric_disposition{name="PowerSupply",successful="true"} 22
ribbon_exporter_metric_disposition{name="SIP ARS",successful="true"} 88
ribbon_exporter_metric_disposition{name="SipStatistic",successful="true"} 88
ribbon_exporter_metric_disposition{name="TrunkGroup",successful="true"} 22
ribbon_exporter_metric_disposition{name="hardDiskUsage",successful="true"} 22
```

## Exporter Metric Duration

```
HELP ribbon_exporter_metric_duration How long metrics took to query and process
# TYPE ribbon_exporter_metric_duration summary
ribbon_exporter_metric_duration{name="CPU",stage="http",quantile="0.5"} 0.064425704
ribbon_exporter_metric_duration{name="CPU",stage="http",quantile="0.9"} 0.068641193
ribbon_exporter_metric_duration{name="CPU",stage="http",quantile="0.99"} 0.069487665
ribbon_exporter_metric_duration_sum{name="CPU",stage="http"} 1.4198724980000001
ribbon_exporter_metric_duration_count{name="CPU",stage="http"} 22
..............
```
