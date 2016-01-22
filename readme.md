# rproxy

A simple Golang _single host_ reverse proxy that implements package httputil's `NewSingleHostReverseProxy` function.

```
go get github.com/etelej/rproxy
```


Usage:
```
rproxy.Run(listenAddr,scheme,remoteHost)
```

Example: 
```
package main

import "github.com/etelej/rproxy"

func main() {
    rproxy.Run(":8081", "dev.etelej.com", "https")
}
```
