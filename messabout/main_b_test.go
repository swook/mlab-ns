package messabout

import (
	"testing"
)

// Taken from http://mlab-ns.appspot.com/admin/sites
var metros = []string{
	"wlg", "vie", "trn", "tpe", "syd", "syd", "svg", "sea",
	"prg", "par", "ord", "nuq", "nuq", "nbo", "mil", "mia",
	"mad", "lju", "lhr", "lga", "lga", "lca", "lba", "lax",
	"jnb", "iad", "hnd", "ham", "dub", "dfw", "atl", "ath",
	"ath", "arn", "ams", "ams", "akl",
}

func Benchmark_ReqWithMetro(b *testing.B) {
	n := len(metros)
	for i := 0; i < b.N; i++ {
		ReqWithMetro(TESTURL, metros[i%n])
	}
}

// IPs generated with http://www.morningstarsecurity.com/research/geoipgen
var ips = []string{
	"170.232.218.106", "142.29.21.202", "147.61.184.234",
	"121.154.27.202", "161.57.238.138", "69.115.227.234",
	"74.23.205.202", "208.216.169.18", "6.127.15.12",
	"180.177.177.210", "121.100.128.106", "75.250.84.42",
	"159.51.120.42", "172.161.59.10", "54.185.63.50",
	"168.33.54.138", "58.22.107.242", "177.231.72.178",
	"186.167.66.210", "190.164.226.242", "81.215.120.138",
	"56.183.208.242", "106.31.169.42", "158.108.168.106",
	"205.240.154.210", "90.56.111.138", "203.19.67.146",
	"177.70.118.146", "168.143.109.74", "143.247.140.42",
	"199.196.57.146", "165.98.82.170", "63.96.164.12",
	"92.59.123.202", "92.113.39.42", "125.40.254.234",
	"3.103.170.18", "82.13.19.234", "17.194.130.160",
	"172.203.35.106", "186.58.204.18", "90.249.74.138",
	"67.252.164.74", "32.121.228.109", "5.49.63.242",
	"173.140.203.210", "99.251.115.170", "175.99.70.50",
	"136.151.255.74", "61.18.158.236", "140.84.4.10",
	"134.18.182.202", "73.65.96.74", "6.180.170.108",
	"22.121.12.82", "70.181.144.42", "9.66.221.10",
	"222.40.131.146", "83.132.91.138", "56.130.53.146",
	"32.175.127.205", "136.44.200.138", "60.221.3.140",
	"101.14.225.42", "164.241.27.234", "100.165.42.106",
	"217.108.222.18", "23.59.169.82", "169.164.182.42",
	"54.131.163.210", "98.24.253.74", "36.53.28.114",
	"15.92.142.234", "119.60.19.74", "35.41.19.146",
	"31.217.18.77", "38.24.154.210", "170.179.63.10",
	"14.234.179.234", "135.247.45.42", "160.254.83.42",
	"149.186.188.202", "116.84.52.10", "14.35.223.199",
	"175.45.166.210", "77.6.93.10", "209.66.194.210",
	"179.165.156.242", "96.161.189.170", "32.229.27.45",
	"136.205.154.170", "83.239.178.74", "66.132.244.170",
	"33.26.182.141", "13.176.101.106", "80.93.193.23",
	"17.140.231.64", "172.53.9.90", "215.191.169.18",
	"156.112.212.106",
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
