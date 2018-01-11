/* GetHeadInfo
 */
package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
/*
The first point to note is the almost excessive amount of error checking that is going on.
This is normal for networking programs:
	the opportunities for failure are substantially greater than for standalone programs.
	Hardware may fail on the client, the server, or on any of the routers and switches in the middle.
	communication may be blocked by a firewall;
	timeouts may occur due to network load;
	the server may crash while the client is talking to it.

The following checks are performed:
	There may be syntax errors in the address specified
	The attempt to connect to the remote service may fail. For example, the service requested might not be running, or there may be no such host connected to the network
	Although a connection has been established, writes to the service might fail if the connection has died suddenly, or the network times out
	Similarly, the reads might fail

 */