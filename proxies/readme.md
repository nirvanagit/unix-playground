## Proxies

### Reverse
coming into the server, think of it as a load balancer

### Forward
Obvious from the name, this is where clients are sending requests to servers

In Reverse proxy, servers receive requests from the proxy, and in Forward proxy clients send requests via a proxy.

## Setup
Open 3 terminal windows
```
# first
go run main.go

# second
go run proxy.go

# third
➜  unix-playground git:(master) ✗ curl -Lv --proxy http://localhost:8080 localhost:4080 -H "Foo: master"
* Rebuilt URL to: localhost:4080/
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 8080 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET http://localhost:4080/ HTTP/1.1
> Host: localhost:4080
> User-Agent: curl/7.54.0
> Accept: */*
> Proxy-Connection: Keep-Alive
> Foo: master
>
< HTTP/1.1 200 OK
< Content-Length: 22
< Content-Type: text/plain; charset=utf-8
< Date: Sat, 11 Jan 2020 04:58:11 GMT
<
* Connection #0 to host localhost left intact
Welcome to my website!
```

logs from main.go
```
➜  server git:(master) ✗ go run main.go
2020/01/10 20:54:50 HTTP/1.1 - GET
2020/01/10 20:54:50 Headers: map[Accept:[*/*] Accept-Encoding:[gzip] Foo:[master] Proxy-Connection:[Keep-Alive] User-Agent:[curl/7.54.0] X-Forwarded-For:[127.0.0.1]]
2020/01/10 20:58:11 HTTP/1.1 - GET
2020/01/10 20:58:11 Headers: map[Accept:[*/*] Accept-Encoding:[gzip] Foo:[master] Proxy-Connection:[Keep-Alive] User-Agent:[curl/7.54.0] X-Forwarded-For:[127.0.0.1]]
```

logs from proxy.go
```
2020/01/10 20:58:11 127.0.0.1:53763   GET   http://localhost:4080/
2020/01/10 20:58:11 127.0.0.1:53763   200 OK
```

## References
- https://gist.github.com/yowu/f7dc34bd4736a65ff28d
