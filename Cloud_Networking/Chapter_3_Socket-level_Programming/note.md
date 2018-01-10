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