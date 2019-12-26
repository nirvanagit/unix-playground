package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// unixgram -> SOCK_DGRAM
	conn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{"/tmp/unixgram", "unixgram"})
	if err != nil {
		panic(err)
	}
	defer os.Remove("/tmp/unixgram")

	for {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(buf[:n]))
		conn.Close()
	}
}
