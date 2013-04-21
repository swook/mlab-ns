package messabout

import (
	"testing"
)

// Taken from http://mlab-ns.appspot.com/admin/sites
var metros = []string{
	"wlg", "vie", "trn", "tpe", "syd", "svg", "sea", "prg",
	"par", "ord", "nuq", "nbo", "mil", "mia", "mad", "lju",
	"lhr", "lga", "lca", "lba", "lax", "jnb", "iad", "hnd",
	"ham", "dub", "dfw", "atl", "ath", "arn", "ams", "akl",
}

func Benchmark_ReqWithMetro(b *testing.B) {
	n := len(metros)
	for i := 0; i < b.N; i++ {
		ReqWithMetro(TESTURL, metros[i%n])
	}
}

// IPs generated with http://www.morningstarsecurity.com/research/geoipgen
var ips = []string{
	"155.203.253.3",
	"116.84.46.195",
	"177.231.67.107",
	"130.3.18.131",
	"216.203.204.171",
	"175.45.161.139",
	"211.84.19.75",
	"113.217.174.67",
	"142.243.126.3",
	"99.37.0.227",
	"75.196.179.131",
	"20.47.91.139",
	"171.91.11.227",
	"151.181.226.99",
	"37.66.40.11",
	"174.248.6.43",
	"84.37.72.99",
	"216.96.133.235",
	"18.99.79.121",
	"65.228.29.67",
	"132.150.113.227",
	"205.186.248.43",
	"138.58.213.3",
	"222.40.126.75",
	"46.137.211.139",
	"169.111.149.131",
	"19.57.188.249",
	"163.65.179.163",
	"112.48.207.67",
	"114.227.191.35",
	"87.127.94.99",
	"218.174.143.11",
	"81.108.44.131",
	"115.77.21.227",
	"190.56.78.235",
	"163.119.79.3",
	"79.245.31.3",
	"159.212.69.3",
	"43.130.250.203",
	"27.166.215.139",
	"50.120.95.171",
	"115.179.76.163",
	"99.144.55.163",
	"85.49.145.67",
	"40.102.66.171",
	"101.227.73.99",
	"69.223.29.99",
	"34.145.240.230",
	"208.110.96.11",
	"8.36.35.163",
	"195.161.102.179",
	"96.108.29.3",
	"198.20.46.75",
	"160.196.178.131",
	"62.191.204.165",
	"95.149.163.131",
	"193.243.181.179",
	"205.44.69.11",
	"91.46.224.163",
	"216.150.33.75",
	"11.181.122.195",
	"100.165.37.35",
	"88.139.103.67",
	"51.9.49.203",
	"103.22.156.195",
	"209.66.189.139",
	"200.248.51.139",
	"221.32.117.107",
	"101.67.119.67",
	"40.209.121.107",
	"122.82.81.199",
	"136.98.94.163",
	"120.113.178.67",
	"70.127.239.131",
	"58.112.93.107",
	"177.70.113.75",
	"164.187.123.67",
	"177.16.213.235",
	"223.155.197.235",
	"84.144.127.35",
	"94.137.130.163",
	"176.4.59.11",
	"152.212.235.67",
	"37.226.250.43",
	"137.204.158.67",
	"100.218.192.131",
	"149.240.82.227",
	"117.42.188.67",
	"117.96.103.163",
	"199.88.243.139",
	"136.151.250.3",
	"82.66.170.3",
	"15.92.137.163",
	"117.150.11.3",
	"139.247.160.3",
	"138.114.112.99",
	"188.31.125.43",
	"131.190.4.99",
	"76.208.188.99",
	"112.208.165.99",
}

func Benchmark_ReqWithIP(b *testing.B) {
	n := len(ips)
	for i := 0; i < b.N; i++ {
		ReqWithIP(TESTURL, ips[i%n])
	}
}

// Taken from http://en.wikipedia.org/wiki/ISO_3166-1#Officially_assigned_code_elements
var countrycodes = []string{
	"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ",
	"AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH",
	"BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO",
	"BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF",
	"BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL",
	"CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR",
	"CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM",
	"DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK",
	"FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM",
	"GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU",
	"GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN",
	"HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM",
	"IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI",
	"KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR",
	"LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY",
	"MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX",
	"FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM",
	"NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG",
	"NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA",
	"PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA",
	"RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF",
	"PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC",
	"SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS",
	"SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH",
	"SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO",
	"TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE",
	"GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG",
	"VI", "WF", "EH", "YE", "ZM", "ZW",
}

func Benchmark_ReqWithCountry(b *testing.B) {
	n := len(countrycodes)
	for i := 0; i < b.N; i++ {
		ReqWithCountry(TESTURL, countrycodes[i%n])
	}
}

func Benchmark_youtube(b *testing.B) {
	for i := 0; i < b.N; i++ {
		youtube()
	}
}

func Benchmark_godoc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		godoc()
	}
}
