Running Benchmark
-----------------

To run these simple benchmarks, use

	go test -bench=".*"" -benchtime=15

the parameter, `-benchtime=15` allows for each `Benchmark*` function to run for approx. 15 seconds.


Sample output
-------------

For the following environment:

 - Go version 1.0.3
 - GOARCH = amd64
 - GOOS = linux
 - Location: Seoul, Korea

Benchmark output:

	Benchmark_ReqWithMetro       100         418194790 ns/op
	Benchmark_ReqWithIP          100         488355900 ns/op
	Benchmark_ReqWithCountry     100         405279920 ns/op
	Benchmark_youtube            200         139220475 ns/op
	Benchmark_godoc              200         166715395 ns/op

Note on 'results'
-----------------

`MetroResolver`, `GeoResolver` and `CountryResolver` are used with varying url parameters.

Fresh set of IPs used. GeoResolver results not cached using memcache serverside.

- *MetroResolver* takes approx. 420ms.
- *GeoResolver* takes approx. 490ms.
- *CountryResolver* takes approx. 410ms.

These results are specifically for the case of my laptop which is located, at the time of writing, in Seoul, Korea. The laptop is running Ubuntu Linux 12.10 with KDE 4.10.1.

- *MetroResolver* and *CountryResolver* take around about the same time despite *MetroResolver*'s extra queries in `ResolverBase._get_candidates_from_sites`. This is due to *MetroResolver* memcache-ing queries.
- *GeoResolver* takes 80ms more than *CountryResolver* due to either or both finding min. distance site and Maxmind query. Previous runs have showed higher differences (up to 250ms), so in this case Maxmind queries may be minimal.

To see how slow mlab-ns is, the response time of a [Youtube API](https://developers.google.com/youtube/2.0/reference) query and the response time of GET-ing the [http://godoc.org](http://godoc.org) homepage is done.

Youtube is chosen to compare with a fast API, godoc is chosen as it is written in Go for Google AppEngine.

- *Youtube API* takes approx. 140ms.
- *godoc.org homepage* takes approx. 170ms

It should be noted that godoc.org is a GET request for a large amount of HTML. It can be seen that both requests should take around 140ms.

This is the sort of response time expected for a Go port of mlab-ns.