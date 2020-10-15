package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func httpClient(url string) (*http.Response, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("call to secondary service failed..", err)
	}

	status := strconv.Itoa(resp.StatusCode)
	outboundTotal.WithLabelValues(status, "GET", "httpClient", "localhost", url).Inc()

	responseBytes, _ := ioutil.ReadAll(resp.Body)
	return resp, responseBytes
}
