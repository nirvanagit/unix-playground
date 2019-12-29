## Raw socket programming at link layer

### Receiver
```
➜  link_layer git:(master) ✗ sudo go run receive/receive.go
Password:
45 00 0A 00 28 44 00 00 40 01 00 00 7F 00 00 01 7F 00 00 01 00 00 3F 21 00 00 00 00 C0 DE
45 00 0A 00 ED 43 00 00 40 01 00 00 7F 00 00 01 7F 00 00 01 00 00 3F 21 00 00 00 00 C0 DE
```

### Sender
```
➜  link_layer git:(master) ✗ sudo go run send/send.go
2019/12/28 16:02:44 creating packet
2019/12/28 16:02:44 sending packet: [69 0 30 0 0 0 0 0 64 1 0 0 0 0 0 0 127 0 0 1 8 0 55 33 0 0 0 0 192 222]
2019/12/28 16:02:44 successfully sent packet
```

## References
- [darkcoding.new](https://www.darkcoding.net/software/raw-sockets-in-go-link-layer/)
