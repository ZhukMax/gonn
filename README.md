# gonn
Simple GoLang Feed Forward Neural Network realization.

## Use
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
  "Port": "8080",
  "DataPath": "/home/user/gonn/data"
}
```
