/*
Package rproxy is a simple single host http reverse proxy.

Usage:

   rproxy.Run(listenAddr,scheme,remoteHost)

Example:

   rproxy.Run(":8081","https","dev.etelej.com")

rproxy was written by Peter Etelej <peter@etelej.com>
*/
package rproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Run launches a reverse proxy listening on a specified port
func Run(listenAddr, remoteHost, scheme string) {
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
