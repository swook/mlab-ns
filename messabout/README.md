Running Benchmark
-----------------

To run these simple benchmarks, use

	go test -bench=".*"" -benchtime=15

the parameter, `-benchtime=15` allows for each `Benchmark*` function to run for approx. 15 seconds.


Sample output
-------------

	GOARCH = amd64
	GOOS = linux
	Go version 1.0.3
	Location: Seoul, Korea

	Requesting lookup with metro: wlg
	{"city": "Wellington", "url": "http://ndt.iupui.mlab2.wlg01.measurement-lab.org:7123", "ip": ["103.10.233.24"], "fqdn": "ndt.iupui.mlab2.wlg01.measurement-lab.org", "site": "wlg01", "country": "NZ"}
	Requesting lookup with IP address: 163.7.129.12
	{"city": "Wellington", "url": "http://ndt.iupui.mlab1v4.wlg01.measurement-lab.org:7123", "ip": ["103.10.233.11"], "fqdn": "ndt.iupui.mlab1v4.wlg01.measurement-lab.org", "site": "wlg01", "country": "NZ"}
	Requesting lookup with Country: NL
	{"city": "Amsterdam", "url": "http://ndt.iupui.mlab2.ams01.measurement-lab.org:7123", "ip": ["213.244.128.152"], "fqdn": "ndt.iupui.mlab2.ams01.measurement-lab.org", "site": "ams01", "country": "NL"}
	PASS
	Benchmark_ReqWithMetro        50         492030300 ns/op
	Benchmark_ReqWithIP           50         683522180 ns/op
	Benchmark_ReqWithCountry              50         432231220 ns/op

Note on 'results'
-----------------

`MetroResolver`, `GeoResolver` and `CountryResolver` are used with varying url parameters.

Fresh set of IPs used. GeoResolver results not cached using memcache serverside.

- *MetroResolver* takes approx. 490ms.
- *GeoResolver* takes approx. 680ms.
- *CountryResolver* takes approx. 430ms.

These results are specifically for the case of my laptop which is located, at the time of writing, in Seoul, Korea. The laptop is running Ubuntu Linux 12.10 with KDE 4.10.1.

- *MetroResolver* and *CountryResolver* take around about the same time despite *MetroResolver*'s extra queries in `ResolverBase._get_candidates_from_sites`. This is due to *MetroResolver* memcache-ing queries.
- *GeoResolver* takes 250ms more than *CountryResolver* due to either or both finding min. distance site and Maxmind query. Likely both as delay is long.

To see how slow mlab-ns may be, the response time of a [Youtube API](https://developers.google.com/youtube/2.0/reference) query and the response time of GET-ing the [http://godoc.org](http://godoc.org) homepage is done.