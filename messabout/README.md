Running Benchmark
-----------------

To run these simple benchmarks, use

	go test -bench=".*"" -benchtime=30

the parameter, `-benchtime=30` allows for each `Benchmark*` function to run for approx. 30 seconds.


Sample output (Cold Run)
------------------------

	Requesting lookup with metro: wlg
	{"city": "Wellington", "url": "http://ndt.iupui.mlab1.wlg01.measurement-lab.org:7123", "ip": ["103.10.233.11"], "fqdn": "ndt.iupui.mlab1.wlg01.measurement-lab.org", "site": "wlg01", "country": "NZ"}
	Requesting lookup with IP address: 163.7.129.12
	{"city": "Wellington", "url": "http://ndt.iupui.mlab1v4.wlg01.measurement-lab.org:7123", "ip": ["103.10.233.11"], "fqdn": "ndt.iupui.mlab1v4.wlg01.measurement-lab.org", "site": "wlg01", "country": "NZ"}
	Requesting lookup with Country: NL
	{"city": "Amsterdam", "url": "http://ndt.iupui.mlab3.ams02.measurement-lab.org:7123", "ip": ["72.26.217.103"], "fqdn": "ndt.iupui.mlab3.ams02.measurement-lab.org", "site": "ams02", "country": "NL"}
	PASS
	Benchmark_ReqWithMetro       100         399183710 ns/op
	Benchmark_ReqWithIP          100         461228120 ns/op
	Benchmark_ReqWithCountry             100         390090510 ns/op


Note
----

Requests are memcached on the mlab-ns server.

`MetroResolver`, `GeoResolver` and `CountryResolver` are used with varying url parameters.

Benchmarks are by no means scientific.