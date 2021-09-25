package client

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Job struct {
	Client *http.Client

	ParamsMind map[string][]string
	Params     []string
	DataTransh []string
	AssetID    []string

	GlobalMap       map[string]map[string][]string
	RecipientTransh map[string][]string
	SenderTransh    map[string][]string
	TranshID        map[string][]string
}

func NewClient() *http.Client {
	var j Job

	host := "http://127.0.0.1:8080"
	proxyURL, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		Proxy:               http.ProxyURL(proxyURL),
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 500,
		MaxConnsPerHost:     500,
		ForceAttemptHTTP2:   true,
		DialContext: (&net.Dialer{
			Timeout: time.Duration(time.Duration(15) * time.Second),
		}).DialContext,
		TLSHandshakeTimeout: time.Duration(time.Duration(15) * time.Second),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
	}
	j.Client = &http.Client{Transport: tr}
	return j.Client
}

func NewJob() *Job {
	var j Job
	j.Client = NewClient()
	j.ParamsMind = make(map[string][]string)
	j.Params = nil
	j.DataTransh = nil
	j.AssetID = nil
	j.RecipientTransh = make(map[string][]string)
	j.SenderTransh = make(map[string][]string)
	j.TranshID = make(map[string][]string)
	j.GlobalMap = make(map[string]map[string][]string)
	return &j
}
