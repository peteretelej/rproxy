/*
Package rproxy is a simple single host http reverse proxy.

Usage of rproxy:

Get the package

   go get -u github.com/peteretelej/rproxy

Import into your application:

   import "github.com/peteretelej/rproxy"

Launch rproxy via the Run function:

   rproxy.Run(":8081","https","etelej.com")


rproxy was written by Peter Etelej <peter@etelej.com>
*/
package rproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Run launches a reverse proxy listening on a specified port
func Run(listenAddr, scheme, remoteHost string) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			p := httputil.NewSingleHostReverseProxy(&url.URL{
				Scheme: scheme,
				Host:   remoteHost,
			})
			p.ServeHTTP(w, r)
		})
	http.ListenAndServe(listenAddr, nil)
}
