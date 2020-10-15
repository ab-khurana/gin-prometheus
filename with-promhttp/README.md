
## with-promhttp

A service written in go. 

The endpoints are exposed on port `8080`. 
The service interacts with another service. The requests are monitored with `go-gin-prometheus` library. 

The monitored metrics can be viewed at `/metrics` endpoint. 
