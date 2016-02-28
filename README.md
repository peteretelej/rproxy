rproxy
======
[![GoDoc](https://godoc.org/github.com/etelej/rproxy?status.svg)](https://godoc.org/github.com/etelej/rproxy) [![Build Status](https://api.travis-ci.org/etelej/rproxy.svg?branch=master)](https://travis-ci.org/etelej/rproxy)

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
    rproxy.Run(":8081", "https", "etelej.com")
}
```
