// example of using rules package for defining custom rules.
/*use of:
-NewAlertingRule
-NewGroup
-NewRecordingRule
-group.Eval 
*/
package main

import "https://github.com/prometheus/prometheus/rules" 


rule := NewAlertingRule(
		"HTTPRequestRateLow",
		 prometheus_http_requests_total{job="prometheus",method="get"}/prometheus_http_request_duration_seconds_sum{job="prometheus",method="get"} < 75 ,
		 time.Minute[5m],
		 labels.FromStrings("severity", "{{ if lt $value 75.0 }}critical{{ else }}warning{{ end }}"),
		 nil, nil, true, nil,
	)


rule := NewAlertingRule(
		"PrometheusConfigurationReloadFailure", 
		prometheus_config_last_reload_successful != 1, 
		time.Minute[5m], 
		labels.FromStrings("severity", "warning"),
		labels.FromStrings("Prometheus configuration reload failure (instance {{ $labels.instance }})"), 
		nil,
		false,
		nil
	)

rule := NewAlertingRule(
		"SQLDown", 
		pg_up == 0, 
		time.Minute[5m], 
		labels.FromStrings("severity", "error"),
		labels.FromStrings("sql down (instance {{ $labels.instance }})"), 
		labels.FromStrings( "sql instance is down\n  VALUE = {{ $value }}\n  LABELS = {{ $labels }}"),
		false,
		nil
	)

rule := NewAlertingRule(
		"PrometheusAlertmanagerNotificationFailing", 
		rate(alertmanager_notifications_failed_total[1m]) > 0, 
		time.Minute[5m], 
		labels.FromStrings("severity", "error"),
		labels.FromStrings("Prometheus AlertManager notification failing (instance {{ $labels.instance }})"), 
		labels.FromStrings( "Alertmanager is failing sending notifications\n  VALUE = {{ $value }}\n  LABELS = {{ $labels }}"),
		false,
		nil
	)



// making groups of System Alerts

group := NewGroup(GroupOptions{
		Name:          "System Alerts",
		Interval:      time.Second,
		Rules:         []Rule{
						NewAlertingRule(
							 
							 "HostOutOfMemory",
							 node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100 < 10,
							 time.Minute[5m],
							 labels.FromStrings("severity", "warning"),
							 labels.FromStrings("Host out of memory (instance {{ $labels.instance }})")
							 labels.FromStrings("Node memory is filling up (< 10% left)\n  VALUE = {{ $value }}\n  LABELS: {{ $labels }}", 
							 true, 
							 nil),

						NewAlertingRule(

							 "HostMemoryUnderMemoryPressure",
							 rate(node_vmstat_pgmajfault[1m]) > 1000,
							 time.Minute[5m],
							 labels.FromStrings("severity", "warning"),
							 labels.FromStrings("Host memory under memory pressure (instance {{ $labels.instance }})"
							 labels.FromStrings("The node is under heavy memory pressure. High rate of major page faults\n  VALUE = {{ $value }}\n  LABELS: {{ $labels }}")
							 true,
							 nil,
							 )
						

						NewAlertingRule(

							 "HostUnusualNetworkThroughputIn",
							 sum by (instance) (irate(node_network_receive_bytes_total[2m])) / 1024 / 1024 > 100,
							 time.Minute[5m],
							 labels.FromStrings("severity", "warning"),
							 labels.FromStrings("Host unusual network throughput in (instance {{ $labels.instance }})")
							 labels.FromStrings("Host network interfaces are probably receiving too much data (> 100 MB/s)\n  VALUE = {{ $value }}\n  LABELS: {{ $labels }}")
							 true,
							 nil,
							 )

						// recording rules : precompute frequently needed or computationally expensive expressions 

						NewRecordingRule("record1", (node_filesystem_avail_bytes{mountpoint="/rootfs"}  * 100) / node_filesystem_size_bytes{mountpoint="/rootfs"} < 10, labels.Labels{{Name: "Disk almost full", Value: "v0"}}),
						NewRecordingRule("record2",  predict_linear(node_filesystem_free_bytes{fstype!~"tmpfs"}[1h], 4 * 3600) < 0, labels.Labels{{Name: "disk full in 4hr ", Value: "v1"}}),

						},


		ShouldRestore: true,
		Opts:          opts,
})



// using .Eval function  , Evalauated twice 1 second apart
group.Eval(ctx, time.Unix(0, 0))
group.Eval(ctx, time.Unix(1, 0))





