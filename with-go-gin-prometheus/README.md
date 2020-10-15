
## with-go-gin-prometheus

A service written in go. 

The endpoints are exposed on port `8081`. 
The requests are monitored with `promhttp` with `gin` http web framework. 

The monitored metrics can be viewed at `/metrics` endpoint. 
The metrics is recoding incoming calls made to the service and calls made to external service from within this service. 