package main

import (
	"net"
	"os"
)

func main() {
	t := "unixgram"
	laddr := net.UnixAddr{"/tmp/unixgramclient", t}
	conn, err := net.DialUnix(t, &laddr, &net.UnixAddr{"/tmp/unixgram", t})
	if err != nil {
	}
	defer os.Remove("/tmp/unixdomainclient")

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
	conn.Close()
}
