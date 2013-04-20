package messabout

import (
	"fmt"
	"testing"
)

const TESTURL = "http://mlab-ns.appspot.com/ndt"

func TestReqWithMetro(t *testing.T) {
	metro := "wlg"
	fmt.Printf("Requesting lookup with metro: %s\n", metro)
	fmt.Println(ReqWithMetro(TESTURL, metro))
}

func TestReqWithIP(t *testing.T) {
	ip := "163.7.129.12"
	fmt.Printf("Requesting lookup with IP address: %s\n", ip)
	fmt.Println(ReqWithIP(TESTURL, ip))
}

func TestReqWithCountry(t *testing.T) {
	country := "NL"
	fmt.Printf("Requesting lookup with Country: %s\n", country)
	fmt.Println(ReqWithCountry(TESTURL, country))
}
