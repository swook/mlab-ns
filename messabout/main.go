// Package messabout messes about with making GET requests to mlab-ns
package messabout

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 20 April 2012, 0724 GMT, http://mlab-ns.appspot.com/admin/sliver_tools
// Deployed tool instances count
// - NDT: 113
// - NPAD: 114
// - neubot: 113
// - Broadband: 114
// No. of sites: 37

const (
	DEFARGS = "?policy=geo&format=json"
)

var (
	c *http.Client
)

func init() {
	c = &http.Client{}
}

func reqBase(url string) string {
	resp, err := c.Get(url)
	if err != nil {
		return ""
	}
	out, _ := ioutil.ReadAll(resp.Body)
	return string(out)
}

// ReqWithMetro performs a lookup request with metro specified
func ReqWithMetro(toolURL, metro string) (result string) {
	url := fmt.Sprintf("%s%s&metro=%s", toolURL, DEFARGS, metro)
	return reqBase(url)
}

// ReqWithMetro performs a lookup request with an IPv4 address specified
func ReqWithIP(toolURL, ip string) (result string) {
	url := fmt.Sprintf("%s%s&ip=%s", toolURL, DEFARGS, ip)
	return reqBase(url)
}

// ReqWithMetro performs a lookup request with country code specified
func ReqWithCountry(toolURL, country string) (result string) {
	url := fmt.Sprintf("%s%s&country=%s", toolURL, DEFARGS, country)
	return reqBase(url)
}
