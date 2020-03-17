### Use of Prometheus with Python Flask and MySQL

integrate Prometheus monitoring with web app based on python

prometheus_client library for python 

Flask is web framework , creates http server and I will be able to configure routes.


use of mysqlclient to query MYSQL DB.


1. counter : to capture the amount of times an HTTP endpoint is hit + to capture the amount of times MySQL query is executed.
2. Histogram : to capture latency of HTTP requests and MySQL queries 

#### Requirements 

mysqlclient

flask

prometheus_client
