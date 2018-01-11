package main

import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
)

/*
The net can be any of "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only) and "ip6" IPv6-only).
It will return an appropriate implementation of the Conn interface.
Note that this function takes a string rather than address as raddr argument, so that programs using this can avoid working out the address type first.

Using this function makes minor changes to programs.
For example, the earlier program to get HEAD information from a Web page can be re-written as
 */

func main() {
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
