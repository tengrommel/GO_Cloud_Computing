# Chapter 3 Socket-level Programming
> It deals with host and service addressing,and then considers TCP and UDP.<br>
It shows how to build both servers and clients using the TCP and UDP Go APIs.<br>
It also looks at raw sockets,in case you need to implement your own protocol above IP.

## The TCP/IP stack
>The OSI model was devised using a committee process wherein the standard was set up and then implemented.<br>
Some parts of the OSI standard are obscure, some parts cannot easily be implemented, some parts have not been implemented.

- IP datagrams
>The IP layer provides a connectionless and unreliable delivery system.It considers each datagram independently of the others.<br>
*The IP layer supplies a checksum that includes its own header.The header includes the source and destination addresses.*

*The IP layer handles routing through an Internet. It is also responsible for breaking up large datagrams into smaller ones for transmission and reassembling them at the other end.*

- UDP
>UDP is also connectionless and unreliable. What it adds to IP is a checksum for the contents of the datagram and port numbers

- TCP
>TCP supplies logic to give a reliable connection-oriented protocol above IP. It provides a virtual circuit that two processes can use to communicate.

## IP address type
    type IP []byte
>The type IP is defined as byte slices

## The type IPmask
    type IPMask []byte
> masking type
    
    func IPv4Mask(a, b, c, d byte) IPMask
> There is a function to create a mask from a 4-byte IPv4 address

    func(ip IP) DefaultMask() IPMask
> Alternatively,there is a method of IP which returns the default mask

## The type IPAddr
    type IPAddr {
        IP IP
    }
> Many of the other functions and methods in the net package return a pointer to an IPAddr.This is simply a structure containing an IP.

A primary use of this type is to perform DNS lookups on IP host names.
    
    func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
    
## Host lookup
    func LookupHost(name string) (addr []string, err os.Error)
> The function ResolveIPAddr will perform a DNS lookup on a hostname, and return a single IP address.<br>
 However, hosts may have multiple IP addresses, usually from multiple network interface cards.<br> 
 They may also have multiple host names, acting as aliases. 
 
    func LookupCNAME(name string)(cname string, err os.Error)
> One of these addresses will be labelled as the "canonical" host name.

## Ports
    func LookupPort(network, service string) (port int, err os.Error)
>On a Unix system, the commonly used ports are listed in the file /etc/services. 
<br>Go has a function to interrogate this file

## The type TCPAddr
    type TCPAddr struct{
        IP      IP
        Port    int
    }
    
- The function to create a TCPAddr is ResolveTCPAddr

    
    func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
    
**where net is one of "tcp", "tcp4" or "tcp6" and the addr is a string composed of a host name or IP address, followed by the port number after a ":", such as "www.google.com:80" or "127.0.0.1:22". If the address is an IPv6 address, which already has colons in it, then the host part must be enclosed in square brackets, such as "[::1]:23". Another special case is often used for servers, where the host address is zero, so that the TCP address is really just the port name, as in ":80" for an HTTP server.**    

## TCP Sockets
- client
> If you are a client you need an API that will allow you to connect to a service and then to send messages to that service and read replies back from the service.
- server 
> If you are a server,you need to be able to bind to a port and listen at it. When a message comes in you need to be able to read it and write back to the client.

The net.TCPConn is the Go type which allows full duplex communication between the client and the server. Two major methods of interest are

    func (c *TCPConn) Write(b []byte) (n int, err os.Error)
    func (c *TCPConn) Read(b []byte) (n int, err os.Error)                                                       
**A TCPConn is used by both a client and a server to read and write messages.**

### TCP client
1. Once a client has established a TCP address for a service, it "dials" the service.
2. If successful, the dial returns TCPConn for communication.
3. The client and the server exchange message on this.
4. Typically a client writes a request to the server using the TCPConn, and reads a response from the TCPConn.
5. This continues until either (or both) sides close the connection.
6. A TCP connection is established by the client using the function
    
    
    func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
>laddr is the local address which is usually set to nil <br>
 raddr is the remote address of the service <br>
 the net string is one of "tcp4", "tcp6" or "tcp" depending on whether you want a TCPv4 connection, a TCPv6 connection or don't care.
 
 ## A Daytime server
 > About the simplest service that we can build is the daytime service.
 
**A server register itself on a port, and listens on that port.
 Then it blocks on an "accept" operation, waiting for clients to connect.**

The relevant calls are:
    
    func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
    func (l *TCPListener) Accept() (c Conn, err os.Error)
    
*The argument net can be set to one of the strings "tcp", "tcp4" or "tcp6". <br>
The IP address should be set to zero if you want to listen on all network interfaces, <br>
or to the IP address of a single network interface if you only want to listen on that interface.<br>
If the port is set to zero, then the O/S will choose a port for you. Otherwise you can choose your own.<br> 
Note that on a Unix system, you cannot listen on a port below 1024 unless you are the system supervisor, root, <br>
and ports below 128 are standardized by the IETF. The example program chooses port 1200 for no particular reason.<br>
The TCP address is given as ":1200" - all interfaces, port 1200.* 

## Controlling TCP connections

- Timeout
> The server may wish to timeout a client if it does not respond quickly enough i.e. does not write a request to the server in time.<br>
This should be a long period(server minutes), because the user may be taking their time.<br>
Conversely, the client may want to timeout the server (after a much shorter time).
    
    func (c *TCPConn) SetTimeout(nsec int64) os.Error
before any reads or writes on the socket.

- Staying alive
> A client may wish to stay connected to a server even if it has nothing to send. It can use

    func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
    
## UDP Datagrams
> In a connectionless protocol each message contains information about its origin destination.<br>
There is no "session" established using a long-lived socket.<br>
UDP clients and servers make use of datagrams, which are individual messages containing source and destination information.<br>
There is no state maintained by these messages, unless the client or server does so. <br>
The messages are not guaranteed to arrive, or may arrive out of order.

The major difference between TCP and UDP handling for Go is how to deal with packets arriving from possibly multiple clients

    func ResolveUDPAddr(net, addr, string) (*UDPAddr, os.Error)
    func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
    func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
    func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error)
    func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
    
## Server listening on multiple sockets
> A server may be attempting to listen to multiple clients not just on one port, but on many.<br>
**In this case it has to use some sort of polling mechanism between the ports.**

1. In C, the select() call lets the kernel do this work. The call takes a number of file descriptors.
2. The process is suspended. When I/O is ready on one of these, a wakeup is done, and the process can continue. This is cheaper than busy polling. 
3. In Go, accomplish the same by using a different goroutine for each port. A thread will become runnable when the lower-level select() discovers that I/O is ready for this thread.

## The types Conn, PacketConn and Listener
>The type Conn is an interface and both TCPConn and UDPConn implement this interface. To a large extent you can deal with this interface rather than the two types.
**Instead of separate dial functions for TCP and UDP, you can use a single function**
    
    func Dial(net, laddr, raddr string) (c Conn, err os.Error)
    
