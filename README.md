# GoNN
Simple GoLang Feed Forward Neural Network realization.

## Use
### Examlpe of simple ANN http-service 
```go
package main

import (
    "fmt"
    "gonn/api_interface"
)

func main() {
    fmt.Print("Start")
    config := gonn.GetConfig("config.json")
    api_interface.Server(config)
}
```
config.json
```json
{
  "Port": "56128",
  "DataPath": "/home/user/gonn/data"
}
```
### Http requests to the service
#### make new Network
```console
curl --data-urlencode "title=Test Net&" http://localhost:56128/new
```
