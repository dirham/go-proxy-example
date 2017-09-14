package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Transport struct{}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Host = req.URL.Host
	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	rp := httputil.NewSingleHostReverseProxy(&url.URL{
		// Scheme: "https" //for https,
		Scheme: "http",
		Host:   "kartu.dtc.co.id",
	})
	rp.Transport = &Transport{}

	http.Handle("/", rp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
