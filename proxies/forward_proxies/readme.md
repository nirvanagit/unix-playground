## Creating a forward proxy

What we need
1. A docker container
2. 2 network namespaces
3. privoxy binary
4. curl
5. a simple server, which returns the remote address of a request

### Setup

```
➜  unix-playground git:(master) ✗ docker run -it --privileged=true ubuntu:14.04 sh
# apt-get install privoxy
```

Checkout how to create network namespaces in - [network namespaces](https://github.com/nirvanagit/unix-playground/commit/f042b37010f12ab44b73011b67f7b30d6f53ea9c)

Build the proxies/forward_proxies/server go program for linux, and copy inside the container
```
➜  unix-playground git:(master) ✗ GOOS="linux" go build ./proxies/forward_proxies/server/...
➜  unix-playground git:(master) ✗ docker cp server 83b5dc45118e:/home/
```

Open 3 shell sessions to the container

Start privoxy in namespace1 (ip - 192.168.1.11)
```
# sed -i -e '/^listen-address/s/localhost/0.0.0.0/' \
    -e '/^accept-intercepted-requests/s/0/1/' \
    -e '/^enforce-blocks/s/0/1/' \
    -e '/^#debug/s/#//' /etc/privoxy/config

# ip netns exec namespace1 privoxy --no-daemon /etc/privoxy/config
```

Start server in namespace2 (ip - 192.168.1.12)
```
# ip netns exec namespace2 /home//server
```

Check if a direct call to server works
```
# ifconfig | grep br1 -A 4
br1       Link encap:Ethernet  HWaddr 36:48:03:5a:c4:b9
          inet addr:192.168.1.10  Bcast:192.168.1.255  Mask:255.255.255.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:205 errors:0 dropped:0 overruns:0 frame:0
          TX packets:103 errors:0 dropped:0 overruns:0 carrier:0
#
#
#
# curl 192.168.1.12:4080/whoyare
{"addr":"192.168.1.10:59788"}#
```
Where `192.168.1.10` is the ip of the container

Call server via proxy
```
# http_proxy="http://192.168.1.11:8118" curl -v http://192.168.1.12:4080/whoyare
* Hostname was NOT found in DNS cache
*   Trying 192.168.1.11...
* Connected to 192.168.1.11 (192.168.1.11) port 8118 (#0)
> GET http://192.168.1.12:4080/whoyare HTTP/1.1
> User-Agent: curl/7.35.0
> Host: 192.168.1.12:4080
> Accept: */*
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 12 Jan 2020 18:50:57 GMT
< Content-Length: 29
< Proxy-Connection: keep-alive
<
* Connection #0 to host 192.168.1.11 left intact
{"addr":"192.168.1.11:50942"}#
```
Where ip `192.168.1.11` is the ip of namesepace1, where privoxy is running.

## References
- inspiration from - [golang-foward-proxy](https://gianarb.it/blog/golang-forwarding-proxy)
- [privoxy configruation](https://github.com/gianarb/dockerfile/blob/master/privoxy/Dockerfile)
