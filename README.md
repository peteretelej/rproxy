rproxy
======
[![GoDoc](https://godoc.org/github.com/peteretelej/rproxy?status.svg)](https://godoc.org/github.com/peteretelej/rproxy) [![Build Status](https://api.travis-ci.org/peteretelej/rproxy.svg?branch=master)](https://travis-ci.org/peteretelej/rproxy)

A simple Golang _single host_ reverse proxy that implements package httputil's `NewSingleHostReverseProxy` function.

```
go get github.com/peteretelej/rproxy
```


Usage:
```
rproxy.Run(listenAddr,scheme,remoteHost)
```

Example: 
```
package main

import "github.com/peteretelej/rproxy"

func main() {
    rproxy.Run(":8081", "https", "etelej.com")
}
```
