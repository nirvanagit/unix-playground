## Network namespaces

From [man7.org](http://man7.org/linux/man-pages/man7/network_namespaces.7.html)

Network namespaces provide isolation of the system resources
       associated with networking: network devices, IPv4 and IPv6 protocol
       stacks, IP routing tables, firewall rules, the /proc/net directory
       (which is a symbolic link to /proc/PID/net), the /sys/class/net
       directory, various files under /proc/sys/net, port numbers (sockets),
       and so on.  In addition, network namespaces isolate the UNIX domain
       abstract socket namespace (see unix(7)).

       A physical network device can live in exactly one network namespace.
       When a network namespace is freed (i.e., when the last process in the
       namespace terminates), its physical network devices are moved back to
       the initial network namespace (not to the parent of the process).

       A virtual network (veth(4)) device pair provides a pipe-like
       abstraction that can be used to create tunnels between network
       namespaces, and can be used to create a bridge to a physical network
       device in another namespace.  When a namespace is freed, the veth(4)
       devices that it contains are destroyed.

       Use of network namespaces requires a kernel that is configured with
       the CONFIG_NET_NS option.


## Exercise

In a docker environment 
1. Create two network namespaces and assign ip addresses to each.
2. Connect them to default namespace using a virtual ethernet device.
3. Start a netcat server on each of them on port `8080`.
4. Make a call to both the IPs in different shell sessions.

Setup
```
docker run -it --privileged=true ubuntu:14.04 sh
# apt-get update
..
..
# apt-get install tree
```

### Setting up two network namespaces

Create namespace
```
# ip netns add namespace1
# ip netns add namespace2

# tree /var/run/netns
/var/run/netns
|-- namespace1
`-- namespace2

0 directories, 2 files
```

Setup veth pairs
```
# ip link add veth1 type veth peer name br-veth1
# ip link add veth2 type veth peer name br-veth2

# # Associate with corresponding namespace
# ip link set veth1 netns namespace1
# ip link set veth2 netns namespace2
```

Add ip address to each veth interface
```
# ip netns exec namespace1 ip addr add 192.168.1.11/24 dev veth1
# ip netns exec namespace2 ip addr add 192.168.1.12/24 dev veth2
```

Create bridge device
```
# ip link add name br1 type bridge
# ip link set br1 up

# ip link | grep br1
8: br1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UKNOWN mode DEFAULT group default qlen 1000
```

Connect bridge with veth pair to the bridge device
```
# ip link set br-veth1 up
# ip link set br-veth2 up

# # Set veth from namespace up
# ip netns exec namespace1 ip link set veth1 up
# ip netns exec namespace2 ip link set veth2 up

# # Setup br-veth* interface to the bridge by setting the bridge device as their master
# ip link set br-veth1 master br1
# ip link set br-veth2 master br2

# # check the bridge is the master of the two interfaces we created
# bridge link show br1
```

Give an address to the bridge device
```
# ip addr add 192.168.1.10/24 brd + dev br1
```

Check connectivity from the default namespace to namespace1 and namespace2

```
# ping 192.168.1.11
PING 192.168.1.11 (192.168.1.11) 56(84) bytes of data.
64 bytes from 192.168.1.11: icmp_seq=1 ttl=64 time=0.050 ms
64 bytes from 192.168.1.11: icmp_seq=2 ttl=64 time=0.061 ms

# ping 192.168.1.12
PING 192.168.1.12 (192.168.1.12) 56(84) bytes of data.
64 bytes from 192.168.1.12: icmp_seq=1 ttl=64 time=0.050 ms
64 bytes from 192.168.1.12: icmp_seq=2 ttl=64 time=0.061 ms
```

### Starting a server in both namespaces using netcat

Open 2 terminal windows and exec into docker container
```
# # window 1
# ip netns exec namespace1 nc -l 8080

# # window 2
# ip netns exec namespace2 nc -l 8080
```

Open 2 terminal windows and exec into docker container
```
# # window 1
# telnet 192.168.1.11 8080
Trying 192.168.1.11...
Connected to 192.168.1.11.
Escape character is '^]'.

# # window 2
# telnet 192.168.1.12 8080
Trying 192.168.1.12...
Connected to 192.168.1.12.
Escape character is '^]'.
```

Open another terminal window

```
# # run netstat on each namespace
# ip netns exec namespace1 netstat -tunapl
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 192.168.1.11:8080       192.168.1.10:58188      ESTABLISHED 197/nc
#
#
# ip netns exec namespace2 netstat -tunapl
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 192.168.1.12:8080       192.168.1.10:68188      ESTABLISHED 191/nc
#
```

## Conclusion
With network namespace we created an isolated network stack, which has its own ipv4, ipv6, ports etc, which is proven by the
fact that we were able to run a server listening the same port `8080` in each namespace, and were able to reach it using the ip
assigned to them.

## References
- Awesome blog on creating network namespaces (all of the work shown above, except running it in docker, and running netcat client/servers) - [ops.tips](https://ops.tips/blog/using-network-namespaces-and-bridge-to-isolate-servers/)
- [network namespaces](https://ops.tips/blog/using-network-namespaces-and-bridge-to-isolate-servers/)
