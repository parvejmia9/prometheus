# prometheus

```shell
go run main.go start -p 8181
localhost:8181/ping
localhost:8181/metrics
```
Run this command to execute 10 pings to the server. Then check the counter
increases on the metrics url. 
```shell
./ping_test.sh
```