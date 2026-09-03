### Rate Limiting HTTP Requests in Go based on IP address

code experiment based on [article from Alex Pliutau](https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g)

#### let's try first without the middleware:

spoiler : we will get 1'000'000 http 200 responses at a speed of 112'580 Requests/sec !

```go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
```

```bash
cgil@vortex:~/go/src/github.com/lao-tseu-is-alive/golang-learning/http$ hey -z 15s -n 1000  http://localhost:8080/ -H 'Host: localhost' -H 'Connection: keep-alive'

Summary:

  Total:	15.0008 secs
  Slowest:	0.0219 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	 112579.9495
  
  Total data:	37153380 bytes
  Size/request:	37 bytes

Response time histogram:

  0.000 [1]	|
  0.002 [991087]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [6279]	|
  0.007 [1688]	|
  0.009 [573]	|
  0.011 [262]	|
  0.013 [74]	|
  0.015 [26]	|
  0.017 [8]	|
  0.020 [0]	|
  0.022 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0004 secs
  75% in 0.0005 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0021 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0219 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0040 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0163 secs
  resp wait:	0.0006 secs, 0.0000 secs, 0.0174 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0136 secs

Status code distribution:
  [200]	1000000 responses

cgil@vortex:~/go/src/github.com/lao-tseu-is-alive/golang-learning/http$ vegeta attack -duration=10s -rate=100 -targets=./vegeta.conf | vegeta report
Requests      [total, rate, throughput]  1000, 100.10, 100.10
Duration      [total, attack, wait]      9.990378795s, 9.990005443s, 373.352µs
Latencies     [mean, 50, 95, 99, max]    391.758µs, 387.63µs, 450.649µs, 529.578µs, 1.469842ms
Bytes In      [total, mean]              22000, 22.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:1000  
Error Set:

```

#### now let's see the how it goes with the limitMiddleware:

spoiler : we will get "only" 700 http 200 responses and
all the other 999'300 responses will be rejected http 429 "Too Many Requests"

_that's a cool way to limit some aggressive users ip !_ 

```go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	if err := http.ListenAndServe(":8080", limitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
```

```bash
cgil@vortex:~/go/src/github.com/lao-tseu-is-alive/golang-learning/http$ hey -z 15s -n 1000  http://localhost:8080/ -H 'Host: localhost' -H 'Connection: keep-alive'

Summary:
  Total:	15.0007 secs
  Slowest:	0.0171 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	107946.1363
  
  Total data:	29150570 bytes
  Size/request:	29 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [984570]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [10593]	|
  0.005 [3056]	|
  0.007 [984]	|
  0.009 [449]	|
  0.010 [260]	|
  0.012 [70]	|
  0.014 [9]	|
  0.015 [6]	|
  0.017 [2]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0004 secs
  75% in 0.0005 secs
  90% in 0.0007 secs
  95% in 0.0010 secs
  99% in 0.0023 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0171 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0028 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0121 secs
  resp wait:	0.0006 secs, 0.0000 secs, 0.0170 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0124 secs

Status code distribution:
  [200]	700 responses
  [429]	999300 responses

cgil@vortex:~/go/src/github.com/lao-tseu-is-alive/golang-learning/http$ vegeta attack -duration=10s -rate=100 -targets=./vegeta.conf | vegeta report

Requests      [total, rate, throughput]  1000, 100.10, 1.40
Duration      [total, attack, wait]      9.990351298s, 9.989950481s, 400.817µs
Latencies     [mean, 50, 95, 99, max]    412.489µs, 409.185µs, 472.987µs, 529.58µs, 2.17872ms
Bytes In      [total, mean]              18056, 18.06
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    1.40%
Status Codes  [code:count]               200:14  429:986  
Error Set:
429 Too Many Requests

```

