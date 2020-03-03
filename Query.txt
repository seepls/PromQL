http_requests_total{job="apiserver", handler="/api/comments"}[5m]

http_requests_total{job=~".*server"} 
# all jobs end with server 

http_requests_total{status!~"4.."}
# select all HTTP code except 4.X.X

rate(http_requests_total[5m])[30m:1m]

max_over_time(deriv(rate(distance_covered_total[5s])[30s:5s])[10m:])

sum by (job) (
  rate(http_requests_total[5m])
)

(instance_memory_limit_bytes - instance_memory_usage_bytes) / 1024 / 1024


sum by (app, proc) (
  instance_memory_limit_bytes - instance_memory_usage_bytes
) / 1024 / 1024

