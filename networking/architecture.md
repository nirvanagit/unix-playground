## Terminoligies

### Router
It is a networking device which forwards data packets between computer networks. It is connected to two or more data lines
from different IP Networks. When data packet comes in one one of the lines, the router reads the network address information
in the packet header to determine the ultimate destination [wikipedia](https://en.wikipedia.org/wiki/Router_(computing))

### Switches
A switch is a layer 2 device, it can learn MAC addresses, uses full-duplex (can receive and send data at the same time).
Because it can learn MAC addresses, it saves a huge amount of bandwidth.

How saving MAC addresses saves bandwidth?
Before switches were introduced, LANs used Hubs. When a packet arrived in a hub, it was sent to all the connected devices. Now lets say there
are 4 devices A, B, C and D connected to a hub, if A sends a packet to C, it goes via a hub, which sends it to all three, B, C and D. And whichever
device accepts is accepted as the receiver.

Devices are connected to a switch via a port in the switch. When device A, which is connected to the switch on port 1, sends a
packet to the switch, switch stores its MAC address and its corresponding port.

| PORT | MAC Address  |
|------|--------------|
|  1   |   02:11:AA   |

This table is temporary, and has a default refresh of 300 seconds.

At this point, the switch does not have any information about the MAC Address of C. To find this out Switches uses __Flooding__.
In this process a switch will send the packets to all the connected devices on the ports other than from where the packet originated, and then once a
device receives the packets it responds with its MAC address which gets stored in the switch's MAC Address table.

ARP table resolves IP to a MAC Address (layer 3 logical address to a layer 2 physical address)

MAC Address to Physical Port table is used by a Switch.

OSI model

| Pneumonic | Layers          |   Number   |     Data      |
|-----------|-----------------|------------|---------------|
|  Aunty    |   Application   |     7      |     Data      |
|  Pinky    |   Presentation  |     6      |     Data      |
|  Simple   |   Session       |     5      |     Data      |
|  Thank    |   Transport     |     4      |     Segments  |
|  Not      |   Network       |     3      |     Packets   |
|  Do       |   Data          |     2      |     Frames    |
|  Please   |   Physical      |     1      |     Bits      |

### How does a packet travel from one computer to the other?

Packet needs to travel from Host A to Host D.

Host A in network 11.11.11.0/24
Host D in network 22.22.22.0/24

1. Host A has data for Host D, which is on a different network, A and D are connected like A->Switch_X->Router->Switch_Y->D.
   - Host A already knows Host D's IP Address.
   - Host A creates a L3 Header.
   - Host A needs to learn its Default Gateway's MAC Address.
2. Host A sends an ARP Request for 11.11.11.1, which is the IP of the default gateway (router).
3. Switch_X Receives the Frame
   - Switch X Learns MAC Address mapping on port 2 (on which A is connected)
   - Switch X sends an ARP request with IP 11.11.11.1, floods frame out to all the ports, this way all the devices connected to SwitchX will get a request,
     and all devices but one will accept it and send a response back.
4. All other devices accept the router discard the frame.
5. Router Accepts the frame when it see that the request was for its own IP address. It processes the request
   - Router learns Host A's ARP mapping
6. Router generates an ARP response.
7. Switch X Receives the frame
   - Switch X Learns MAC adress mapping on port 3 (which the route connected to)
   - Switch X forwards the frame out to port 2 (where A is connected)
8. How A receives the ARP response
   - Host A learns the Routers ARP mapping, which means it has its MAC address.
9. Host A creates a L2 Header, which has the SRC MAC address as that of A, and DST MAC Address as that of the router.
   - It also creates an L3 Header, which has the SRC IP Address of itself, and the DST IP Address of Host B.
   - L3's Headers purpose is to get the data end-to-end, from one host to the other.
   - L2's Headers purpose is to get the data from hop to hop.
10. Frame arrives on Switch X.
   - It already knows the mapping for port 2, hence wil refresh.
   - Sends the packet via port 3 to the router.
11. Router receives the packet and strips the L2 Header. It strips it because the requirement of this header was to get
    data from Host A to the router.
12. Router checks the L3 Header and consults its routing table, and realizes that the IP address matches the network on one of its interfaces.
    - 22.22.22.0/24 network exists on eth2 interface. (router has two interfaces, eth1, which is connected to 11.11.11.0/24 and eth2 connected to 22.22.22.0/24)
    - Router needs to learn the MAC address for 22.22.22.40 (host B)
13. Router sends an ARP request for 22.22.22.40.
14. Switch Y receives the frame.
    - Switch Y Learns the MAC Address mapping on port 4 (where router is connected to)
    - Switch Y Floods the frame out to all the ports, except from where it received it.
15. All hosts except the B discard the frame.
16. Host D receives the ARP request, and hence processes the ARP request.
    - Host D learns the Routers ARP mapping, which means it stores the router's MAC Address.
17. Host D generates an ARP response and sends a UNICAST to the router with the router's MAC address in its L2 Header.
18. Switch Y receives the frame on port 5 (which is where B is connected to it)
    - Switch Y Learns the MAC address mapping on port 5.
    - Switch Y Forwards frame out to port 4, which is where the router is connected to the switch)
19. Router receives the ARP response, and learns Host D's MAC address.
    - Now router has the L2 header details (MAC address of D) to send the packet to D.
20. Router creates a L2 Header with SRC MAC as its port connected to 22.22.22.0/24 network, and the
    DST MAC address as that of D.
21. Switch Y receives the Frame on port 4 (connected to the router)
    - Switch Y already knows mapping for port 4, which it refreshes.
    - Switch Y forwards the frame out of port 5, which is where Host D is connected to.
22. Host D receives the Packet.
    - Host D strips the Layer 2 Header, as its not needed anymore.
    - Host D strips L3 header, as the packet has already reached the destination.

Response from Host D
23. Host D responds with Data for Host A
    - Host D builds a L3 Header with Host A's IP address. Host D can compare the IP address of A and determine that
      it belongs to a foreign IP.
    - Host D needs to send the packet to its Default Gateway.
    - Host D knows its Default Gateway's MAC address.
24. Host D creates L2 header and sends the packet to Swtich Y on port 5 (which is where Host D is connected to it).
25. Switch Y receives the frame.
    - Switch Y already knows the mapping of Port 5 (connected to host D) so it just refreshes its table.
    - Switch Y forwards the frame out from port 4 (which is connected to the default gateway, i.e. the router)
26. Router receives the packet and strips the L2 Header (as its no longer required because it has already reached
    the Dest MAC address it had)
27. Router consults the routing table
    - 11.11.11.0/24 networks exists on eth1 interface of the router.
28. Router already has an entry in its ARP table for Host A, hence it creates an L2 Header with DST MAC address as
    that of A.
29. Switch X receives the frame.
    - Switch X already knows mapping for port 3.
    - Switch X forwards the frame out of port 2 (which is connected to Host A)
30. Host A receives the packet
    - It strips L2 Header
    - It strips L3 Header
    - Host A receives and responds to D.

### What are VLANs (Virtual Local Area Networks)?

VLANs allow you to break one Physical Switch into multiple Virtual Switches.

If a network is constructed like

Host A -> Switch A -> Router 1 -> Switch B -> Router 2 -> Switch C -> Host D
With the help of VLANS we could convert the above to

Host A -> VLAN A -> Router 1 -> VLAN B -> Router 2 -> VLAN C -> Host D

Ports from one Switch, lets say Switch Z are given to each VLAN, which creates each of the Virtual Switch A B and C.

### What are hubs (old technology replaced by switches)

- A hubs job is to connect devices in a network together.

- Its a layer 1 device

- A sends a packet for C. On receiving the packet, the HUB will send the data to all the other devices other than A,
that means it will send it to B C and D. Host B and D will discard it when they realize it was not meant for them.

- Hubs use half duplex, which means that they cannot send and receive data at the same time without causing a data collision, which
  will cause the data to be corrupt and will need to be sent again.

- waste bandwidth by sending data to all the connected devices.

```
        B
        | 
        |
A ---- HUB ---- C
        | 
        |
        D
```

### What are bridges

- Its a layer 2 device, which means it can learn the source and the destination MAC address of the devices connected to it.
- It segments LANS into smaller networks.
- It has 2 collision domains, that means data can be sent or received on each section at the same time.
- They are not used and are replaced by switches.

```
         B
         | 
         |
A ---- BRIDGE ---- C
         | 
         |
         D
```

### What are switches

Can learn which port is connected to which device and stores its MAC Address.

- Layer 2 Device
- Full Duplex - it can send and receive data at the same time. each port has its own collision domain
- Saves Bandwidth as it saves MAC addresses. This prevents other hosts from stealing data not meant for them

```
        B
        | 
        |
A ---- SWT ---- C
        | 
        |
        D
```

### What are routers
- Layer 3 devices to connect networks
- Usually has 2 ports (depends)

  ``` 
                                         -------
                                   -----/       \----
  A ---- SWITCH ----- ROUTER ---- /      INTERNET    \
                                  \                  /
                                   ------------------
                                  
  ```

### What is a NAT?
Why are they needed?
IPv4 addresses are limited, and we can not have more than ~4 BILLION IP Addresses.

How do they solve it?
Was to create private addresses. These addresses are
- 10.0.0.0 - 10.255.255.255
- 172.16.0.0 - 172.31.255.255
- 192.168.0.0 - 192.168.255.255

These addresses can only be used in an internal network, for instance your home or your office.

```
  (192.168.1.5)                           -------
  C ----                            -----/       \----
        \                          /                   \ 
  (192.168.1.6)                   /                     \ 
  A ---- SWITCH ----- ROUTER ---- |        INTERNET     |
        /                         \                     / 
  B ---                            \                   /
  (192.168.1.7)                     ------------------
                                    
```

Types of NATs
- PAT
- Dynamic
- Static


## References
- https://learningnetwork.cisco.com/thread/6816
- https://www.globalknowledge.com/us-en/resources/resource-library/articles/how-switches-work/
- https://www.youtube.com/watch?v=rYodcvhh7b8
- https://www.youtube.com/watch?v=MmwF1oHOvmg
- https://www.youtube.com/watch?v=qij5qpHcbBk
