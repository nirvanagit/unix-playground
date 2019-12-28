## Reading packets

Note: gopacket/main.go is configured to read packets passed on the loopback interface.

Run a simple golang server
```
go run server/main.go
```

In another terminal run
```
while true; do curl localhost:9000; echo "---"; done;
```

Open another terminal and run
```
# go run gopacket/main.go
```

Sample output from gopacket/main.go on intercepting a packet on the wire
```
➜  unix-playground git:(master) ✗ go run sockets/raw_sockets/examples/gopacket/main.go
packet dump: -- FULL PACKET DATA (88 bytes) ------------------------------------
00000000  1e 00 00 00 60 0f 79 64  00 2c 06 40 00 00 00 00  |....`.yd.,.@....|
00000010  00 00 00 00 00 00 00 00  00 00 00 01 00 00 00 00  |................|
00000020  00 00 00 00 00 00 00 00  00 00 00 01 cc 8f 00 50  |...............P|
00000030  a2 80 fa 35 00 00 00 00  b0 02 ff ff 00 34 00 00  |...5.........4..|
00000040  02 04 3f c4 01 03 03 06  01 01 08 0a 30 e9 90 b1  |..?.........0...|
00000050  00 00 00 00 04 02 00 00                           |........|
--- Layer 1 ---
Loopback        {Contents=[30, 0, 0, 0] Payload=[..84..] Family=IPv6}
00000000  1e 00 00 00                                       |....|
--- Layer 2 ---
IPv6    {Contents=[..40..] Payload=[..44..] Version=6 TrafficClass=0 FlowLabel=1014116 Length=44 NextHeader=TCP HopLimit=64 SrcIP=::1 DstIP=::1 HopByHop=nil}
00000000  60 0f 79 64 00 2c 06 40  00 00 00 00 00 00 00 00  |`.yd.,.@........|
00000010  00 00 00 00 00 00 00 01  00 00 00 00 00 00 00 00  |................|
00000020  00 00 00 00 00 00 00 01                           |........|
--- Layer 3 ---
TCP     {Contents=[..44..] Payload=[] SrcPort=52367 DstPort=80(http) Seq=2726361653 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=52 Urgent=0 Options=[..8..] Padding=[0]}
00000000  cc 8f 00 50 a2 80 fa 35  00 00 00 00 b0 02 ff ff  |...P...5........|
00000010  00 34 00 00 02 04 3f c4  01 03 03 06 01 01 08 0a  |.4....?.........|
00000020  30 e9 90 b1 00 00 00 00  04 02 00 00              |0...........|

packet string: PACKET: 88 bytes, wire length 88 cap length 88 @ 2019-12-28 15:13:03.654648 -0800 PST
- Layer 1 (04 bytes) = Loopback {Contents=[30, 0, 0, 0] Payload=[..84..] Family=IPv6}
- Layer 2 (40 bytes) = IPv6     {Contents=[..40..] Payload=[..44..] Version=6 TrafficClass=0 FlowLabel=1014116 Length=44 NextHeader=TCP HopLimit=64 SrcIP=::1 DstIP=::1 HopByHop=nil}
- Layer 3 (44 bytes) = TCP      {Contents=[..44..] Payload=[] SrcPort=52367 DstPort=80(http) Seq=2726361653 Ack=0 DataOffset=11 FIN=false SYN=true RST=false PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=65535 Checksum=52 Urgent=0 Options=[..8..] Padding=[0]}
```
